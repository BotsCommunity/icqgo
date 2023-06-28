package scene

import (
	"github.com/botscommunity/icqgo/API"
	"github.com/botscommunity/icqgo/update"
)

type Scenes struct {
	Message_         Message
	EditMessage_     EditMessage
	DeleteMessage_   DeleteMessage
	CallbackQuery_   CallbackQuery
	PinMessage_      PinMessage
	UnpinMessage_    UnpinMessage
	ChatMembers_     ChatMembers
	LeftChatMembers_ LeftChatMembers
}

type Message func(bot *API.Bot, message update.Message)
type EditMessage func(bot *API.Bot, message update.EditMessage)
type DeleteMessage func(bot *API.Bot, message update.DeleteMessage)
type CallbackQuery func(bot *API.Bot, message update.CallbackQuery)
type PinMessage func(bot *API.Bot, message update.PinMessage)
type UnpinMessage func(bot *API.Bot, message update.UnpinMessage)
type ChatMembers func(bot *API.Bot, message update.ChatMembers)
type LeftChatMembers func(bot *API.Bot, message update.LeftChatMembers)

func Create() *Scenes {
	return &Scenes{}
}
