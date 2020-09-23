package bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"strings"
)

func (b *Bot) SpeakHandler(m *tb.Message) {
	r := b.d.Get(m.Text)
	if r != "" {
		b.bot.Send(m.Chat, r)
		return
	}
}

func (b *Bot) SpeakAddHandler(m *tb.Message) {
	if m.ReplyTo == nil {
		b.bot.Send(m.Chat, "Вы должны обратиться на то сообщение, к чему будет ответ")
		return
	}
	text := strings.Split(m.Text, " ")
	if len(text) <= 1 {
		return
	}
	b.d.Add(m.ReplyTo.Text, strings.Join(text[1:], " "))

	b.bot.Send(m.Chat, "Шарманка обновлена")
}
