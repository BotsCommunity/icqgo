package tests

import (
	"os"
	"testing"

	"github.com/botscommunity/icqgo/API"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func TestDeleteMembers(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	bot := API.Create(os.Getenv("TOKEN"), logger)

	t.Run("DeleteMembers", func(t *testing.T) {
		bot.DeleteMembers("687964914@chat.agent", API.DeleteChatMembers{
			{
				Sn: "1010794571",
			},
		})
	})
}
