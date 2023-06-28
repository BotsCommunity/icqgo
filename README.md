<img src="preview.png">

<p align="center">Flexible and high-performance library for iCQ Bot API</p>

<div align="center">
	<a href="https://pkg.go.dev/github.com/botscommunity/icqgo">
		<img src="https://img.shields.io/static/v1?label=go&message=reference&color=00add8&logo=go" />
	</a>
	<a href="http://www.opensource.org/licenses/MIT">
		<img src="https://img.shields.io/badge/license-MIT-lime.svg" />
	</a>
</div>

<h2 align="center">Installation</h2>

```bash
go get github.com/botscommunity/icqgo
```

<h2 align="center">Getting Started</h2>

```go
package main

import (
	"github.com/botscommunity/icqgo/API"
	"github.com/botscommunity/icqgo/longpoll"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()

	var (
        bot = API.Create(os.Getenv("TOKEN"))
	    scenes = scene.Create()
    )

    scenes.Message(func(bot *API.Bot, message update.Message) {
		bot.SendMessage(message.Chat.ID, "echo message: " + message.Text)
	})

	longpoll.Create(bot, scenes).Listen()
}
```