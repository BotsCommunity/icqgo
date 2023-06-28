package tests

import (
	"os"
	"testing"

	"github.com/botscommunity/icqgo/API"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func TestChatControl(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	bot := API.Create(os.Getenv("TOKEN"), logger)

	t.Run("SetChatTitle", func(t *testing.T) {
		bot.SetChatTitle("687964914@chat.agent", "Title")
	})

	t.Run("SetChatAbout", func(t *testing.T) {
		bot.SetChatAbout("687964914@chat.agent", "About")
	})

	t.Run("SetChatRules", func(t *testing.T) {
		bot.SetChatRules("687964914@chat.agent", "Rules:\n1. Do not spam")
	})
}
