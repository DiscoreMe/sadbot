package bot

import (
	"encoding/json"
	"fmt"
	"github.com/DiscoreMe/sadbot/weather"
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"strings"
)

const weatherText = `
🌤 %s 🌧

🌡 Температура: %.1f° (макс %.1f°)
💧 Влажность: %d г/м³
💪🏼 Давление: %d мм рт.
💨 Ветер:
         Скорость: %d м/с
         Направление: %d°
`

func (b *Bot) WeatherHandler(m *tb.Message) {
	command := strings.Split(m.Text, " ")
	command = command[1:]
	country := strings.Join(command, " ")

	var data weather.Data

	w, err := b.c.WeatherGet(country)
	if err != nil {
		// todo: log error
		logrus.Errorln(err)
	}

	if len(w) != 0 {
		err := json.Unmarshal(w, &data)
		if err != nil {
			logrus.Errorln(err)
		}
	} else {
		d, err := b.w.WeatherByLocation(country)
		if err != nil {
			b.bot.Send(m.Chat, "Тут два исхода: такого города нет либо бот опять сломался 😕")
			return
		}
		data = *d

		bd, err := json.Marshal(d)
		if err != nil {
			logrus.Errorln(err)
		}
		if err := b.c.WeatherSet(country, bd); err != nil {
			logrus.Errorln(err)
		}
	}

	b.bot.Send(m.Chat, fmt.Sprintf(weatherText, strings.Title(strings.ToLower(country)), data.Main.Temp, data.Main.TempMax, data.Main.Humidity, data.Main.Pressure, data.Wind.Speed, data.Wind.Deg))
}
