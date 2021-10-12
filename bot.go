package botty

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	botApi          *tgbotapi.BotAPI
	defaultHandler  Handler
	messageHandlers map[string]Handler
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

		for msg, handler := range b.messageHandlers {
			if strings.ToLower(update.Message.Text) == msg {
				go handler(MessageIn(*update.Message), messageChan)
			} else {
				go b.defaultHandler(MessageIn(*update.Message), messageChan)
			}
			message := <-messageChan
			if _, err := b.botApi.Send(message.Config); err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *Bot) AddMessageHandler(msgReq string, handler Handler) {
	b.messageHandlers[strings.ToLower(msgReq)] = handler
}

func (b *Bot) DefaultHandler(handler Handler) {
	b.defaultHandler = handler
}

func NewBot(apiKey string) (Bot, error) {
	botApi, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return Bot{}, err
	}

	return Bot{
		messageHandlers: make(map[string]Handler),
		botApi:          botApi,
	}, nil
}
