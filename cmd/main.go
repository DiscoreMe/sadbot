package main

import (
	"github.com/DiscoreMe/sadbot/bot"
	"github.com/DiscoreMe/sadbot/cache"
	"github.com/DiscoreMe/sadbot/config"
	"github.com/DiscoreMe/sadbot/dict"
	"github.com/DiscoreMe/sadbot/weather"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.New()

	d := dict.NewDict()
	w := weather.NewWeather(cfg.OpenWeatherToken)
	c, err := cache.NewCache("127.0.0.1:6379")
	if err != nil {
		logrus.Fatalln("redis connect", err)
	}
	b, err := bot.NewBot(bot.BotSettings{
		Token:   cfg.Token,
		Weather: w,
		Cache:   c,
		Dict:    d,
	})
	if err != nil {
		logrus.Fatalln("initial bot", err)
	}

	b.Listen()
}
