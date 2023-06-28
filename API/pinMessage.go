package API

import "net/url"

type PinMessage struct {
	ChatID    string
	MessageID string
}

func (bot *Bot) PinMessage(properties ...any) (banned ResponseInfo) {
	var (
		params = make(url.Values)
		chatID bool
	)

	for _, property := range properties {
		switch property := property.(type) {
		case PinMessage:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
				chatID = true
			}
			if property.MessageID != "" {
				params.Set("msgId", property.MessageID)
			}

		case ChatID:
			params.Set("chatId", string(property))
			chatID = true
		case MessageID:
			params.Set("msgId", string(property))

		case string:
			if !chatID {
				params.Set("chatId", property)
				chatID = true
			} else {
				params.Set("msgId", property)
			}
		}
	}

	err := bot.Call("chats/pinMessage", params, &banned)
	if err != nil {
		return banned
	}

	return banned
}

type UnpinMessage struct {
	ChatID    string
	MessageID string
}

func (bot *Bot) UnpinMessage(properties ...any) (banned ResponseInfo) {
	var (
		params = make(url.Values)
		chatID bool
	)

	for _, property := range properties {
		switch property := property.(type) {
		case UnpinMessage:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
				chatID = true
			}
			if property.MessageID != "" {
				params.Set("msgId", property.MessageID)
			}

		case ChatID:
			params.Set("chatId", string(property))
			chatID = true
		case MessageID:
			params.Set("msgId", string(property))

		case string:
			if !chatID {
				params.Set("chatId", property)
				chatID = true
			} else {
				params.Set("msgId", property)
			}
		}
	}

	err := bot.Call("chats/unpinMessage", params, &banned)
	if err != nil {
		return banned
	}

	return banned
}
