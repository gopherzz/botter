# botty
A library for quick and easy creation of telegram bots.
Use simple handlers for messages as you would do in a library net/http.

# Example Message Handle
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
 
 # Default Handler Example
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