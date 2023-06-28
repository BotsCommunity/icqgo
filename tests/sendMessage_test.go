package tests

import (
	"os"
	"testing"

	"github.com/botscommunity/icqgo/API"
	"github.com/botscommunity/icqgo/keyboard"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func TestSendMessage(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	bot := API.Create(os.Getenv("TOKEN"), logger)

	t.Run("SendMessage", func(t *testing.T) {
		buttons := keyboard.CreateInline()

		buttons.Add(keyboard.InlineButton{
			Text:         "Button #1",
			CallbackData: "RXhhbXBsZSBjYWxsYmFjayBkYXRhCg==",
			Style:        "primary",
		}).Row()
		buttons.Add(keyboard.InlineButton{
			Text:  "Button #1",
			URL:   "https://api.github.com",
			Style: "primary",
		})

		bot.SendMessage("687964914@chat.agent", "Hello, world!", buttons)
	})

	t.Run("SendFormattedMessage", func(t *testing.T) {
		bot.SendMessage("687964914@chat.agent", `<b>bold</b>, <strong>bold</strong>
		<i>italic</i>, <em>italic</em>
		<u>underline</u>, <ins>underline</ins>
		<s>strikethrough</s>, <strike>strikethrough</strike>, <del>strikethrough</del>
		<a href="http://www.example.com/">inline URL</a>
		<a>inline mention of a user</a>
		<code>inline fixed-width code</code>
		<pre>pre-formatted fixed-width code block</pre>
		<pre><code class="python">pre-formatted fixed-width code block written in the Python programming language</code></pre>
		Ordered list:
		<ol>
		<li>First element</li>
		<li>Second element</li>
		</ol>
		Unordered list:
		<ul>
		<li>First element</li>
		<li>Second element</li>
		</ul>
		Quote:
		<blockquote>
		Begin of quote.
		End of quote.
		</blockquote>`, API.ParseMode("HTML"))
	})

	t.Run("SendCustomTypedMessage", func(t *testing.T) {
		buttons := keyboard.CreateInline()

		buttons.Add(keyboard.InlineButton{
			Text:         "Button #1",
			CallbackData: "RXhhbXBsZSBjYWxsYmFjayBkYXRhCg==",
			Style:        "primary",
		}).Row()
		buttons.Add(keyboard.InlineButton{
			Text:  "Button #1",
			URL:   "https://api.github.com",
			Style: "primary",
		})

		bot.SendMessage("687964914@chat.agent", "CustomTypesMessage", API.InlineKeyboard(buttons))
	})
}
