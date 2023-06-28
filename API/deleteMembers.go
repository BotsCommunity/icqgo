package API

import (
	"encoding/json"
	"net/url"
)

type DeleteMembers struct {
	ChatID  string
	Members []DeleteChatMember
}

type DeleteChatMembers []DeleteChatMember
type DeleteChatMember struct {
	Sn string `json:"sn,omitempty"`
}

func (bot *Bot) DeleteMembers(properties ...any) (members ResponseInfo) {
	var params = make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case DeleteMembers:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
			}
			if property.Members != nil {
				sn, _ := json.Marshal(property.Members)
				params.Add("members", string(sn))
			}

		case ChatID:
			params.Set("chatId", string(property))
		case DeleteChatMembers:
			sn, _ := json.Marshal(property)
			params.Add("members", string(sn))

		case string:
			params.Set("chatId", property)
		}
	}

	err := bot.Call("chats/members/delete", params, &members)
	if err != nil {
		return members
	}

	return members
}
