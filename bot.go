package botty

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	handlers map[string]Handler
	botApi   *tgbotapi.BotAPI
}

func (b *Bot) Start(timeout int) (err error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout

	updates, err := b.botApi.GetUpdatesChan(u)
	messageChan := make(chan Message)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		for msg, handler := range b.handlers {
			if update.Message.Text == msg {
				go handler(update, messageChan)
				if _, err := b.botApi.Send(<-messageChan); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (b *Bot) AddMessageHandler(msgReq string, handler Handler) {
	b.handlers[msgReq] = handler
}

func NewBot(apiKey string) (Bot, error) {
	botApi, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return Bot{}, err
	}

	return Bot{
		handlers: make(map[string]Handler),
		botApi:   botApi,
	}, nil
}
