package API

import (
	"net/url"
)

type SendActions struct {
	ChatID  string
	Actions []Action
}

const (
	Typing  Action = "typing"
	Looking Action = "looking"
)

func (bot *Bot) SendActions(properties ...any) (sent ResponseInfo) {
	params := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case SendActions:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
			}
			if property.Actions != nil {
				for _, action := range property.Actions {
					params.Add("actions", string(action))
				}
			}

		case ChatID:
			params.Set("chatId", string(property))
		case []Action:
			for _, action := range property {
				params.Add("actions", string(action))
			}
		case Actions:
			for _, action := range property {
				params.Add("actions", string(action))
			}

		case string:
			params.Set("chatId", property)
		case []string:
			for _, action := range property {
				params.Add("actions", action)
			}
		}
	}

	err := bot.Call("chats/sendActions", params, &sent)
	if err != nil {
		return sent
	}

	return sent
}
