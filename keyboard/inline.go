package keyboard

import "encoding/json"

type InlineKeyboard struct {
	Buttons InlineButtons
}

type InlineButtons [][]InlineButton
type InlineButton struct {
	Text         string `json:"text,omitempty"`
	URL          string `json:"url,omitempty"`
	CallbackData string `json:"callbackData,omitempty"`
	Style        string `json:"style,omitempty"`
}

func CreateInline() *InlineKeyboard {
	keyboard := &InlineKeyboard{}
	keyboard.Buttons = make([][]InlineButton, 0, 2)
	return keyboard
}

func (keyboard *InlineKeyboard) Add(button InlineButton) *InlineKeyboard {
	count := len(keyboard.Buttons)

	if count == 0 {
		keyboard.Buttons = append(keyboard.Buttons, []InlineButton{})
		keyboard.Buttons[0] = append(keyboard.Buttons[0], button)
	} else {
		keyboard.Buttons[count-1] = append(keyboard.Buttons[count-1], button)
	}

	return keyboard
}

func (keyboard *InlineKeyboard) Row() *InlineKeyboard {
	keyboard.Buttons = append(keyboard.Buttons, make([]InlineButton, 0))
	return keyboard
}

// JSON 🍷🧣🍁 turns an object into a JSON format string
func (keyboard *InlineKeyboard) JSON() string {
	if data, err := json.Marshal(keyboard.Buttons); err == nil {
		return string(data)
	}
	return ""
}

// JSON 🍷🧣🍁 turns an object into a JSON format string
func (buttons *InlineButtons) JSON() string {
	if data, err := json.Marshal(buttons); err == nil {
		return string(data)
	}
	return ""
}
