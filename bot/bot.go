package bot

import (
	"bytes"
	"github.com/DiscoreMe/sadbot/cache"
	"github.com/DiscoreMe/sadbot/calculator"
	"github.com/DiscoreMe/sadbot/config"
	"github.com/DiscoreMe/sadbot/dict"
	"github.com/DiscoreMe/sadbot/weather"
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

type Bot struct {
	bot    *tb.Bot
	w      *weather.Weather
	c      *cache.Cache
	d      *dict.Dict
	calc   *calculator.Cal
	screen config.ScreenConfig

	screenCh chan *screenTask
}

type BotSettings struct {
	Token        string
	Weather      *weather.Weather
	Cache        *cache.Cache
	Dict         *dict.Dict
	Calc         *calculator.Cal
	ScreenConfig config.ScreenConfig
}

func NewBot(settings BotSettings) (*Bot, error) {
	b, err := tb.NewBot(tb.Settings{
		Token:  settings.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		bot:      b,
		c:        settings.Cache,
		d:        settings.Dict,
		w:        settings.Weather,
		calc:     settings.Calc,
		screenCh: make(chan *screenTask),
		screen:   settings.ScreenConfig,
	}
	bot.setup()

	go bot.listenScreenCh()

	return bot, nil
}

func (b *Bot) Listen() {
	b.bot.Start()
}

func (b *Bot) setup() {
	b.bot.Handle(".hello", func(m *tb.Message) {
		b.bot.Send(m.Sender, "Hello World!")
	})
	b.bot.Handle(tb.OnText, b.CmdHandler)
}

func (b *Bot) CmdHandler(m *tb.Message) {
	if m.Text == "" {
		return
	}
	if m.Text[0] != '.' {
		b.SpeakHandler(m)
		return
	}
	if utf8.RuneCountInString(m.Text) <= 1 {
		return
	}

	cmd := strings.Split(m.Text, " ")[0][1:]
	switch cmd {
	case "погода":
		b.WeatherHandler(m)
	case "адик":
		b.SpeakAddHandler(m)
	case "эбауте":
		b.about(m)
	case "кл":
		b.CalcHandler(m)
	case "скрин":
		b.ScreenHandler(m)
	default:
		b.SpeakHandler(m)
	}
}

func (b *Bot) about(m *tb.Message) {
	b.bot.Send(m.Chat, "Исходный код:\nhttps://github.com/DiscoreMe/sadbot")
}

type screenTask struct {
	m    *tb.Message
	b    bytes.Buffer
	url  string
	done bool
}

func (b *Bot) listenScreenCh() {
	for t := range b.screenCh {
		err := func() error {
			client := &http.Client{
				Timeout: 10 * time.Second,
			}

			var u url.URL
			q := u.Query()
			q.Set("key", b.screen.Token)
			q.Set("url", t.url)
			u.RawQuery = q.Encode()

			resp, err := client.Post(b.screen.URL, "application/x-www-form-urlencoded", strings.NewReader(q.Encode()))
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			_, err = io.Copy(&t.b, resp.Body)
			if err != nil {
				return err
			}

			fff, _ := os.Create("test.png")
			io.Copy(fff, resp.Body)
			fff.Close()

			t.done = true

			return nil
		}()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"type": "screen",
				"url":  t.url,
			})
			continue
		}
		b.ScreenResultHandler(t)
	}
}
