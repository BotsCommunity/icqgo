package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/botscommunity/icqgo/API"
	"github.com/botscommunity/icqgo/longpoll"
	"github.com/botscommunity/icqgo/scene"
	"github.com/botscommunity/icqgo/update"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func TestLongPoll(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	t.Run("LongPoll", func(t *testing.T) {
		var (
			bot    = API.Create(os.Getenv("TOKEN"), logger)
			scenes = scene.Create()
		)

		scenes.Message(func(bot *API.Bot, message update.Message) {
			bot.SendMessage(message.Chat.ID, message.Text)
		}).EditMessage(func(bot *API.Bot, message update.EditMessage) {
			fmt.Println("EditMessage is", message)
		}).DeleteMessage(func(bot *API.Bot, message update.DeleteMessage) {
			fmt.Println("DeleteMessage is", message)
		}).Callback(func(bot *API.Bot, query update.CallbackQuery) {
			bot.AnswerCallback(query.ID, "OK", "https://api.github.com/")
		})

		scenes.PinMessage(func(bot *API.Bot, message update.PinMessage) {
			fmt.Println("PinMessage is", message)
		}).UnpinMessage(func(bot *API.Bot, message update.UnpinMessage) {
			fmt.Println("UnpinMessage is", message)
		}).ChatMembers(func(bot *API.Bot, members update.ChatMembers) {
			fmt.Println("ChatMembers is", members)
		}).LeftChatMembers(func(bot *API.Bot, members update.LeftChatMembers) {
			fmt.Println("LeftChatMembers is", members)
		})

		longpoll.Create(bot, scenes).Listen()
	})
}
