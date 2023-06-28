package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/botscommunity/icqgo/API"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func TestGetChat(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	bot := API.Create(os.Getenv("TOKEN"), logger)

	t.Run("GetChat", func(t *testing.T) {
		chat := bot.GetChat("687964914@chat.agent")
		fmt.Printf("Chat is %+v\n", chat)
	})

	t.Run("GetAdmins", func(t *testing.T) {
		admins := bot.GetAdmins("687964914@chat.agent")
		fmt.Printf("Chat Admins is %+v\n", admins)
	})

	t.Run("GetMembers", func(t *testing.T) {
		members := bot.GetMembers("687964914@chat.agent")
		fmt.Printf("Chat Members is %+v\n", members)
	})

	t.Run("GetBlockedUsers", func(t *testing.T) {
		users := bot.GetBlockedUsers("687964914@chat.agent")
		fmt.Printf("Blocked users is %+v\n", users)
	})

	t.Run("GetPendingUsers", func(t *testing.T) {
		users := bot.GetPendingUsers("687964914@chat.agent")
		fmt.Printf("Pending users %+v\n", users)
	})
}
