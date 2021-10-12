package botty

type Handler func(MessageIn, chan<- Message)
type DefaultHandler func(MessageIn, chan<- Message)
