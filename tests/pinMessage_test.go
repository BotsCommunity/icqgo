package tests

import (
	"os"
	"testing"

	"github.com/botscommunity/icqgo/API"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func TestPinMessage(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	bot := API.Create(os.Getenv("TOKEN"), logger)

	t.Run("PinMessage", func(t *testing.T) {
		bot.PinMessage("687964914@chat.agent", "7249691494687178796")
	})

	t.Run("UnpinMessage", func(t *testing.T) {
		bot.UnpinMessage("687964914@chat.agent", "7249691494687178796")
	})
}
