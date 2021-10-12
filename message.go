package botty

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Message tgbotapi.MessageConfig

func NewMessage(update tgbotapi.Update) Message {
	return Message(tgbotapi.NewMessage(update.Message.Chat.ID, ""))
}

func (m *Message) Text(text string) Message {
	m.Text = text
	return *m
}

func (m *Message) Keyboard(kb tgbotapi.InlineKeyboardMarkup) Message {
	m.ReplyMarkup = kb
	return *m
}
