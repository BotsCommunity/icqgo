package API

import (
	"fmt"
	"net/url"
)

type AnswerCallback struct {
	QueryID   string
	Text      string
	URL       string
	ShowAlert bool
}

func (bot *Bot) AnswerCallback(properties ...any) (answered ResponseInfo) {
	var (
		params    = make(url.Values)
		installed = map[string]bool{}
	)

	for _, property := range properties {
		switch property := property.(type) {
		case AnswerCallback:
			if property.QueryID != "" {
				params.Set("queryId", property.QueryID)
				installed["queryID"] = true
			}
			if property.Text != "" {
				params.Set("text", string(property.Text))
				installed["text"] = true
			}
			if property.URL != "" {
				params.Set("url", property.URL)
			}

			if property.ShowAlert {
				params.Set("showAlert", fmt.Sprintf("%t", property.ShowAlert))
			}

		case QueryID:
			params.Set("queryId", string(property))
			installed["queryID"] = true
		case Text:
			params.Set("text", string(property))
			installed["text"] = true
		case URL:
			params.Set("url", string(property))
		case ShowAlert:
			params.Set("showAlert", fmt.Sprintf("%t", property))

		case string:
			if !installed["queryID"] {
				params.Set("queryId", property)
				installed["queryID"] = true
			} else if !installed["text"] {
				params.Set("text", property)
				installed["text"] = true
			} else {
				params.Set("url", property)
			}
		case bool:
			params.Set("showAlert", fmt.Sprintf("%t", property))
		}
	}

	err := bot.Call("messages/answerCallbackQuery", params, &answered)
	if err != nil {
		return answered
	}

	return answered
}
