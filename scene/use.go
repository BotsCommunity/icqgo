package scene

import (
	"encoding/json"
	"fmt"

	"github.com/botscommunity/icqgo/API"
	"github.com/botscommunity/icqgo/update"
)

func (scenes Scenes) Use(bot *API.Bot, event update.Update) {
	if bot.Logger != nil {
		event, _ := json.Marshal(event)
		bot.Logger.Info("Updates response is ", string(event))
	}

	switch event.Type {
	case "newMessage":
		message := update.Message{}
		if err := json.Unmarshal(event.Payload, &message); err != nil {
			fmt.Println(err)
		}

		if scenes.Message_ != nil {
			scenes.Message_(bot, message)
		}
	case "editedMessage":
		message := update.EditMessage{}
		if err := json.Unmarshal(event.Payload, &message); err != nil {
			fmt.Println(err)
		}

		if scenes.EditMessage_ != nil {
			scenes.EditMessage_(bot, message)
		}
	case "deletedMessage":
		message := update.DeleteMessage{}
		if err := json.Unmarshal(event.Payload, &message); err != nil {
			fmt.Println(err)
		}

		if scenes.DeleteMessage_ != nil {
			scenes.DeleteMessage_(bot, message)
		}
	case "callbackQuery":
		query := update.CallbackQuery{}
		if err := json.Unmarshal(event.Payload, &query); err != nil {
			fmt.Println(err)
		}

		if scenes.CallbackQuery_ != nil {
			scenes.CallbackQuery_(bot, query)
		}

	case "pinnedMessage":
		message := update.PinMessage{}
		if err := json.Unmarshal(event.Payload, &message); err != nil {
			fmt.Println(err)
		}

		if scenes.PinMessage_ != nil {
			scenes.PinMessage_(bot, message)
		}
	case "unpinnedMessage":
		message := update.UnpinMessage{}
		if err := json.Unmarshal(event.Payload, &message); err != nil {
			fmt.Println(err)
		}

		if scenes.UnpinMessage_ != nil {
			scenes.UnpinMessage_(bot, message)
		}

	case "newChatMembers":
		members := update.ChatMembers{}
		if err := json.Unmarshal(event.Payload, &members); err != nil {
			fmt.Println(err)
		}

		if scenes.ChatMembers_ != nil {
			scenes.ChatMembers_(bot, members)
		}
	case "leftChatMembers":
		members := update.LeftChatMembers{}
		if err := json.Unmarshal(event.Payload, &members); err != nil {
			fmt.Println(err)
		}

		if scenes.LeftChatMembers_ != nil {
			scenes.LeftChatMembers_(bot, members)
		}
	}
}
