package tests

import (
	"os"
	"testing"

	"github.com/botscommunity/icqgo/API"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func TestBan(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	bot := API.Create(os.Getenv("TOKEN"), logger)

	t.Run("Ban", func(t *testing.T) {
		bot.Ban("687964914@chat.agent", "760748135")
	})

	t.Run("Unban", func(t *testing.T) {
		bot.Unban("687964914@chat.agent", "760748135")
	})
}
