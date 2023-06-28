package API

import (
	"fmt"
	"net/url"
)

type Ban struct {
	ChatID         string
	UserID         string
	DeleteMessages bool
}

func (bot *Bot) Ban(properties ...any) (banned ResponseInfo) {
	var (
		params = make(url.Values)
		chatID bool
	)

	for _, property := range properties {
		switch property := property.(type) {
		case Ban:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
				chatID = true
			}
			if property.UserID != "" {
				params.Set("userId", property.UserID)
			}

			if property.DeleteMessages {
				params.Set("delLastMessages", fmt.Sprintf("%t", property.DeleteMessages))
			}

		case ChatID:
			params.Set("chatId", string(property))
			chatID = true
		case UserID:
			params.Set("userId", string(property))
		case DeleteMessages:
			params.Set("delLastMessages", fmt.Sprintf("%t", property))

		case string:
			if !chatID {
				params.Set("chatId", property)
				chatID = true
			} else {
				params.Set("userId", property)
			}
		case bool:
			params.Set("delLastMessages", fmt.Sprintf("%t", property))
		}
	}

	err := bot.Call("chats/blockUser", params, &banned)
	if err != nil {
		return banned
	}

	return banned
}

type Unban struct {
	ChatID string
	UserID string
}

func (bot *Bot) Unban(properties ...any) (banned ResponseInfo) {
	var (
		params = make(url.Values)
		chatID bool
	)

	for _, property := range properties {
		switch property := property.(type) {
		case Unban:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
				chatID = true
			}
			if property.UserID != "" {
				params.Set("userId", property.UserID)
			}
		case ChatID:
			params.Set("chatId", string(property))
			chatID = true
		case UserID:
			params.Set("userId", string(property))
		case string:
			if !chatID {
				params.Set("chatId", property)
				chatID = true
			} else {
				params.Set("userId", property)
			}
		}
	}

	err := bot.Call("chats/unblockUser", params, &banned)
	if err != nil {
		return banned
	}

	return banned
}
