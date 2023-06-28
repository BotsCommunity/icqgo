package tests

import (
	"os"
	"testing"

	"github.com/botscommunity/icqgo/API"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func TestSendVoice(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	bot := API.Create(os.Getenv("TOKEN"), logger)

	t.Run("SendVoice", func(t *testing.T) {
		file, err := os.Open("./attachments/kolya.aac")
		if err != nil {
			panic(err)
		}
		bot.SendVoice("687964914@chat.agent", file)
	})
}
