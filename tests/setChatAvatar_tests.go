package tests

import (
	"os"
	"testing"

	"github.com/botscommunity/icqgo/API"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func TestSetChatAvatar(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	bot := API.Create(os.Getenv("TOKEN"), logger)

	t.Run("SetChatAvatar", func(t *testing.T) {
		file, err := os.Open("./attachments/photo.jpg")
		if err != nil {
			panic(err)
		}
		bot.SetChatAvatar("687964914@chat.agent", file)
	})

}
