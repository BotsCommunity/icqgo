package API

import "net/url"

type GetChat struct {
	ChatID string
}

func (bot *Bot) GetChat(properties ...any) (chat ResponseChat) {
	params := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case GetChat:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
			}

		case ChatID:
			params.Set("chatId", string(property))

		case string:
			params.Set("chatId", property)
		}
	}

	err := bot.Call("chats/getInfo", params, &chat)
	if err != nil {
		return chat
	}

	return chat
}

type GetAdmins struct {
	ChatID string
}

func (bot *Bot) GetAdmins(properties ...any) (admins ResponseChatAdmins) {
	params := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case GetAdmins:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
			}

		case ChatID:
			params.Set("chatId", string(property))

		case string:
			params.Set("chatId", property)
		}
	}

	err := bot.Call("chats/getAdmins", params, &admins)
	if err != nil {
		return admins
	}

	return admins
}

type GetMembers struct {
	ChatID string
}

func (bot *Bot) GetMembers(properties ...any) (members ResponseChatMembers) {
	params := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case GetMembers:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
			}

		case ChatID:
			params.Set("chatId", string(property))

		case string:
			params.Set("chatId", property)
		}
	}

	err := bot.Call("chats/getMembers", params, &members)
	if err != nil {
		return members
	}

	return members
}

type GetBlockedUsers struct {
	ChatID string
}

func (bot *Bot) GetBlockedUsers(properties ...any) (users ResponseChatUsers) {
	params := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case GetBlockedUsers:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
			}

		case ChatID:
			params.Set("chatId", string(property))

		case string:
			params.Set("chatId", property)
		}
	}

	err := bot.Call("chats/getBlockedUsers", params, &users)
	if err != nil {
		return users
	}

	return users
}

type GetPendingUsers struct {
	ChatID string
}

func (bot *Bot) GetPendingUsers(properties ...any) (users ResponseChatUsers) {
	params := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case GetPendingUsers:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
			}

		case ChatID:
			params.Set("chatId", string(property))

		case string:
			params.Set("chatId", property)
		}
	}

	err := bot.Call("chats/getPendingUsers", params, &users)
	if err != nil {
		return users
	}

	return users
}
