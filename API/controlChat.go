package API

import (
	"net/url"
)

type SetChatTitle struct {
	ChatID string
	Title  string
}

func (bot *Bot) SetChatTitle(properties ...any) (titled ResponseInfo) {
	var (
		params = make(url.Values)
		chatID bool
	)

	for _, property := range properties {
		switch property := property.(type) {
		case SetChatTitle:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
				chatID = true
			}
			if property.Title != "" {
				params.Set("title", property.Title)
			}

		case ChatID:
			params.Set("chatId", string(property))
			chatID = true
		case Title:
			params.Set("title", string(property))

		case string:
			if !chatID {
				params.Set("chatId", property)
				chatID = true
			} else {
				params.Set("title", property)
			}
		}
	}

	err := bot.Call("chats/setTitle", params, &titled)
	if err != nil {
		return titled
	}

	return titled
}

type SetChatAbout struct {
	ChatID string
	About  string
}

func (bot *Bot) SetChatAbout(properties ...any) (titled ResponseInfo) {
	var (
		params = make(url.Values)
		chatID bool
	)

	for _, property := range properties {
		switch property := property.(type) {
		case SetChatAbout:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
				chatID = true
			}
			if property.About != "" {
				params.Set("about", property.About)
			}

		case ChatID:
			params.Set("chatId", string(property))
			chatID = true
		case About:
			params.Set("about", string(property))

		case string:
			if !chatID {
				params.Set("chatId", property)
				chatID = true
			} else {
				params.Set("about", property)
			}
		}
	}

	err := bot.Call("chats/setAbout", params, &titled)
	if err != nil {
		return titled
	}

	return titled
}

type SetChatRules struct {
	ChatID string
	Rules  string
}

func (bot *Bot) SetChatRules(properties ...any) (titled ResponseInfo) {
	var (
		params = make(url.Values)
		chatID bool
	)

	for _, property := range properties {
		switch property := property.(type) {
		case SetChatRules:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
				chatID = true
			}
			if property.Rules != "" {
				params.Set("rules", property.Rules)
			}

		case ChatID:
			params.Set("chatId", string(property))
			chatID = true
		case About:
			params.Set("rules", string(property))

		case string:
			if !chatID {
				params.Set("chatId", property)
				chatID = true
			} else {
				params.Set("rules", property)
			}
		}
	}

	err := bot.Call("chats/setRules", params, &titled)
	if err != nil {
		return titled
	}

	return titled
}
