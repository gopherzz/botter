package botty

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type MessageIn tgbotapi.Message
type Message struct {
	Config tgbotapi.MessageConfig
}

func NewMessage(in MessageIn) Message {
	return Message{Config: tgbotapi.NewMessage(in.Chat.ID, "")}
}

func NewMessageWithText(in MessageIn, text string) Message {
	return Message{Config: tgbotapi.NewMessage(in.Chat.ID, text)}
}

func (m *Message) SetText(text string) Message {
	m.Config.Text = text
	return *m
}

func (m *Message) SetKeyboard(kb tgbotapi.InlineKeyboardMarkup) Message {
	m.Config.ReplyMarkup = kb
	return *m
}
