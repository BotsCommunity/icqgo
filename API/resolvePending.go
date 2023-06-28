package API

import (
	"fmt"
	"net/url"
)

type ResolvePending struct {
	ChatID   string
	UserID   string
	Approve  bool
	Everyone bool
}

func (bot *Bot) ResolvePending(properties ...any) (pending ResponseInfo) {
	var (
		params = make(url.Values)
		chatID bool
	)

	for _, property := range properties {
		switch property := property.(type) {
		case ResolvePending:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
				chatID = true
			}
			if property.UserID != "" {
				params.Set("userId", property.UserID)
			}

			if property.Approve {
				params.Set("approve", fmt.Sprintf("%t", property.Approve))
			}
			if property.Everyone {
				params.Set("everyone", fmt.Sprintf("%t", property.Approve))
			}

		case ChatID:
			params.Set("chatId", string(property))
			chatID = true
		case UserID:
			params.Set("userId", string(property))
		case Approve:
			params.Set("approve", fmt.Sprintf("%t", property))
		case Everyone:
			params.Set("everyone", fmt.Sprintf("%t", property))

		case string:
			if !chatID {
				params.Set("chatId", property)
				chatID = true
			} else {
				params.Set("userId", property)
			}
		case bool:
			params.Set("approve", fmt.Sprintf("%t", property))
		}
	}

	err := bot.Call("chats/resolvePending", params, &pending)
	if err != nil {
		return pending
	}

	return pending
}
