package API

import (
	"net/url"
	"os"

	"github.com/botscommunity/icqgo/keyboard"
)

type SendVoice struct {
	ChatID string
	FileID string

	MessageID  string
	MessagesID []string

	ForwardChatID     string
	ForwardMessageID  string
	ForwardMessagesID []string

	InlineKeyboard any
}

func (bot *Bot) SendVoice(properties ...any) (sent ResponseSentMessage) {
	var (
		params    = make(url.Values)
		file      = File{}
		installed = map[string]bool{}
	)

	for _, property := range properties {
		switch property := property.(type) {
		case SendVoice:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
				installed["chatID"] = true
			}
			if property.FileID != "" {
				params.Set("fileId", property.FileID)
				installed["fileID"] = true
			}

			if property.MessageID != "" {
				params.Add("replyMsgId", property.MessageID)
			}
			if property.MessagesID != nil {
				for _, messageID := range property.MessagesID {
					params.Add("replyMsgId", messageID)
				}
			}

			if property.ForwardChatID != "" {
				params.Set("forwardChatId", property.ForwardChatID)
				installed["forward"] = true
			}
			if property.ForwardMessageID != "" {
				params.Add("forwardMsgId", property.ForwardMessageID)
			}
			if property.ForwardMessagesID != nil {
				for _, messageID := range property.ForwardMessagesID {
					params.Add("forwardMsgId", messageID)
				}
			}

			if property.InlineKeyboard != nil {
				switch property := property.InlineKeyboard.(type) {
				case string:
					params.Set("inlineKeyboardMarkup", property)
				case keyboard.InlineKeyboard:
					params.Set("inlineKeyboardMarkup", property.JSON())
				}
			}

		case *os.File:
			file.Data = property
			installed["fileID"] = true

		case ChatID:
			params.Set("chatId", string(property))
			installed["chatID"] = true
		case FileID:
			params.Set("fileId", string(property))
			installed["fileID"] = true

		case MessageID:
			params.Set("replyMsgId", string(property))
		case MessagesID:
			for _, messageID := range property {
				params.Add("replyMsgId", messageID)
			}
		case ReplyMessageID:
			params.Set("replyMsgId", string(property))
		case ReplyMessagesID:
			for _, messageID := range property {
				params.Add("replyMsgId", messageID)
			}

		case ForwardChatID:
			params.Set("forwardChatId", string(property))
			installed["forward"] = true
		case ForwardMessageID:
			params.Add("forwardMsgId", string(property))
		case ForwardMessagesID:
			for _, messageID := range property {
				params.Add("forwardMsgId", messageID)
			}

		case string:
			if !installed["chatID"] {
				params.Set("chatId", property)
				installed["chatID"] = true
			} else if !installed["fileID"] {
				params.Set("fileId", property)
				installed["fileID"] = true
			} else if installed["forward"] {
				params.Set("forwardMsgId", property)
			} else {
				params.Set("replyMsgId", property)
			}
		case []string:
			if installed["forward"] {
				for _, messageID := range property {
					params.Add("forwardMsgId", messageID)
				}
			} else {
				for _, messageID := range property {
					params.Add("replyMsgId", messageID)
				}
			}

		case *keyboard.InlineKeyboard:
			params.Set("inlineKeyboardMarkup", property.JSON())
		case InlineKeyboard:
			params.Set("inlineKeyboardMarkup", property.Buttons.JSON())
		}
	}

	if file.Data != nil {
		file.Field = "file"
	}

	err := bot.Call("messages/sendVoice", params, &sent, file)
	if err != nil {
		return sent
	}

	return sent
}
