package bot

import (
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"strings"
)

func (b *Bot) ScreenHandler(m *tb.Message) {
	cmd := strings.Split(m.Text, " ")
	if len(cmd) <= 1 {
		return
	}

	if strings.Index(cmd[1], "http") == -1 {
		cmd[1] = "https://" + cmd[1]
	}

	b.screenCh <- &screenTask{
		m:   m,
		url: cmd[1],
	}
}

func (b *Bot) ScreenResultHandler(t *screenTask) {
	if t.done {
		photo := &tb.Photo{File: tb.FromReader(&t.b)}
		photo.Caption = "Сервис скриншота сайтов был любезно предоставлен @HelpfulSenkoBot"
		_, err := b.bot.Reply(t.m, photo)
		if err != nil {
			b.bot.Reply(t.m, "Произошла ошибка при загрузке фото 🙁")
			logrus.Errorln(err)
		}
	} else {
		b.bot.Reply(t.m, "Не удалось заскринить сайт 🙁")
	}
}
