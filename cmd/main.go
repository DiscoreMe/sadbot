package main

import (
	"github.com/DiscoreMe/sadbot/bot"
	"github.com/DiscoreMe/sadbot/cache"
	"github.com/DiscoreMe/sadbot/calculator"
	"github.com/DiscoreMe/sadbot/config"
	"github.com/DiscoreMe/sadbot/dict"
	"github.com/DiscoreMe/sadbot/weather"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cfg := config.New()
	calc := calculator.NewCal()
	d := dict.NewDict()
	w := weather.NewWeather(cfg.OpenWeatherToken)
	c, err := cache.NewCache("127.0.0.1:6379")
	if err != nil {
		logrus.Fatalln("redis connect", err)
	}
	b, err := bot.NewBot(bot.BotSettings{
		Token:        cfg.Token,
		Calc:         calc,
		Weather:      w,
		Cache:        c,
		Dict:         d,
		ScreenConfig: cfg.Screen,
	})
	if err != nil {
		logrus.Fatalln("initial bot", err)
	}

	b.Listen()
}
