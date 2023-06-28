package API

import (
	"encoding/json"
	"net/url"

	"github.com/botscommunity/icqgo/keyboard"
)

type EditMessage struct {
	ChatID    string
	MessageID string
	Text      string

	InlineKeyboard any

	Format    Format
	ParseMode string
}

func (bot *Bot) EditMessage(properties ...any) (info ResponseInfo) {
	var (
		params    = make(url.Values)
		installed = map[string]bool{}
	)

	for _, property := range properties {
		switch property := property.(type) {
		case EditMessage:
			if property.ChatID != "" {
				params.Set("chatId", property.ChatID)
				installed["chatID"] = true
			}
			if property.MessageID != "" {
				params.Add("msgId", property.MessageID)
				installed["messageID"] = true
			}

			if property.Text != "" {
				params.Set("text", property.Text)
			}

			if property.InlineKeyboard != nil {
				switch property := property.InlineKeyboard.(type) {
				case string:
					params.Set("inlineKeyboardMarkup", property)
				case keyboard.InlineKeyboard:
					params.Set("inlineKeyboardMarkup", property.JSON())
				}
			}

			if (Format{}) != property.Format {
				format, _ := json.Marshal(property)
				params.Set("format", string(format))
			}
			if property.ParseMode != "" {
				params.Set("parseMode", property.ParseMode)
			}

		case ChatID:
			params.Set("chatId", string(property))
			installed["chatID"] = true
		case MessageID:
			params.Set("msgId", string(property))
			installed["messageID"] = true
		case Text:
			params.Set("text", string(property))

		case Format:
			format, _ := json.Marshal(property)
			params.Set("format", string(format))
		case ParseMode:
			params.Set("parseMode", string(property))

		case string:
			if !installed["chatID"] {
				params.Set("chatId", property)
				installed["chatID"] = true
			} else if !installed["messageID"] {
				params.Set("msgId", property)
				installed["messageID"] = true
			} else {
				params.Set("text", property)
			}

		case *keyboard.InlineKeyboard:
			params.Set("inlineKeyboardMarkup", property.JSON())
		case InlineKeyboard:
			params.Set("inlineKeyboardMarkup", property.Buttons.JSON())
		}
	}

	err := bot.Call("messages/editText", params, nil)
	if err != nil {
		return info
	}

	return info
}
