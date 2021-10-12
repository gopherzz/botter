# botter
### A library for quick and easy creation of telegram bots based on [telegram-bot-api](github.com/go-telegram-bot-api/telegram-bot-api).
### Use simple handlers for messages as you would do in a [net/http](https://cs.opensource.google/go/go/+/master:src/net/http/) library.

# Examples
## Message Handler Example
 ```go
package main

import (
	"log"

	"github.com/gopherzz/botty"
)

func hello(in botty.MessageIn, out chan<- botty.Message) {
	msg := botty.NewMessage(in)
	msg.Config.Text = "Hello, im botty!"
	out <- msg
}

func main() {
	bot, err := botty.NewBot("YourToken")
	if err != nil {
		log.Fatal(err)
	}
  
	bot.AddMessageHandler("hello", hello)
	bot.Start(60)
}
 ```
 
## Default Handler Example
```go
package main

import (
	"log"

	"github.com/gopherzz/botty"
)

// Simple echo bot, handler
func defHandler(in botty.MessageIn, out chan<- botty.Message) {
	msg := botty.NewMessageWithText(in, in.Text)
	out <- msg
}

func main() {
	bot, err := botty.NewBot("YourToken")
	if err != nil {
		log.Fatal(err)
	}
  
	bot.DefaultHandler(hello)
	bot.Start(60)
}
```
