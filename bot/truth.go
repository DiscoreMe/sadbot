package bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"math/rand"
)

func (b *Bot) TruthHandler(m *tb.Message) {
	r := rand.Int31n(50)
	if r >= 25 {
		b.bot.Reply(m, "Чистая правда")
	} else {
		b.bot.Reply(m, "Наглая ложь")
	}
}
