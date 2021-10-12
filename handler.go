package botty

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type Handler func(tgbotapi.Update, chan<- Message)
