package API

import (
	"net/url"
	"os"
)

type SetChatAvatar struct {
	ChatID string
	File   *os.File
}

func (bot *Bot) SetChatAvatar(properties ...any) (info ResponseInfo) {
	var (
		params = make(url.Values)
		data   = &os.File{}
	)

	for _, property := range properties {
		switch property := property.(type) {
		case SetChatAvatar:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
			}
			if property.File != nil {
				data = property.File
			}

		case *os.File:
			data = property

		case ChatID:
			params.Set("chatId", string(property))

		case string:
			params.Set("chatId", property)
		}
	}

	err := bot.Call("chats/avatar/set", params, nil, File{Field: "image", Data: data})
	if err != nil {
		return info
	}

	return info
}
