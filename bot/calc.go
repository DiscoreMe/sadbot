package bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"strings"
)

func (b *Bot) CalcHandler(m *tb.Message) {
	cmd := strings.Split(m.Text, " ")
	if len(cmd) <= 1 {
		return
	}
	result, err := b.calc.Cal(strings.Join(cmd[1:], ""))
	if err != nil {
		b.bot.Reply(m, "Произошла ошибка в вычислении: "+err.Error())
		return
	}
	b.bot.Reply(m, "Результат: "+result)
}
