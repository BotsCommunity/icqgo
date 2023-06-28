package API

import (
	"net/url"
)

type DeleteMessage struct {
	ChatID string

	MessageID  string
	MessagesID []string
}

func (bot *Bot) DeleteMessage(properties ...any) (info ResponseInfo) {
	var (
		params = make(url.Values)
		chatID bool
	)

	for _, property := range properties {
		switch property := property.(type) {
		case DeleteMessage:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
				chatID = true
			}

			if property.MessageID != "" {
				params.Set("msgId", property.MessageID)
			}
			if property.MessagesID != nil {
				for _, messageID := range property.MessagesID {
					params.Add("msgId", messageID)
				}
			}

		case ChatID:
			params.Set("chatId", string(property))
			chatID = true
		case MessageID:
			params.Set("msgId", string(property))
		case MessagesID:
			for _, messageID := range property {
				params.Add("msgId", messageID)
			}

		case string:
			if !chatID {
				params.Set("chatId", property)
				chatID = true
			} else {
				params.Set("msgId", property)
			}
		case []string:
			for _, messageID := range property {
				params.Add("msgId", messageID)
			}
		}
	}

	err := bot.Call("messages/deleteMessages", params, &info)
	if err != nil {
		return info
	}

	return info
}
