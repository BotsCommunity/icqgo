package scene

func (scenes *Scenes) OnMessage(message Message) *Scenes {
	return scenes.Message(message)
}

func (scenes *Scenes) Message(message Message) *Scenes {
	scenes.Message_ = message
	return scenes
}

func (scenes *Scenes) OnEditMessage(editMessage EditMessage) *Scenes {
	return scenes.EditMessage(editMessage)
}

func (scenes *Scenes) EditMessage(editMessage EditMessage) *Scenes {
	scenes.EditMessage_ = editMessage
	return scenes
}

func (scenes *Scenes) OnDeleteMessage(deleteMessage DeleteMessage) *Scenes {
	return scenes.DeleteMessage(deleteMessage)
}

func (scenes *Scenes) DeleteMessage(deleteMessage DeleteMessage) *Scenes {
	scenes.DeleteMessage_ = deleteMessage
	return scenes
}

func (scenes *Scenes) OnCallback(callbackQuery CallbackQuery) *Scenes {
	return scenes.Callback(callbackQuery)
}

func (scenes *Scenes) Callback(callbackQuery CallbackQuery) *Scenes {
	scenes.CallbackQuery_ = callbackQuery
	return scenes
}

func (scenes *Scenes) OnPinMessage(pinMessage PinMessage) *Scenes {
	return scenes.PinMessage(pinMessage)
}

func (scenes *Scenes) PinMessage(pinMessage PinMessage) *Scenes {
	scenes.PinMessage_ = pinMessage
	return scenes
}

func (scenes *Scenes) OnUnpinMessage(unpinMessage UnpinMessage) *Scenes {
	return scenes.UnpinMessage(unpinMessage)
}

func (scenes *Scenes) UnpinMessage(unpinMessage UnpinMessage) *Scenes {
	scenes.UnpinMessage_ = unpinMessage
	return scenes
}

func (scenes *Scenes) OnChatMembers(chatMembers ChatMembers) *Scenes {
	return scenes.ChatMembers(chatMembers)
}

func (scenes *Scenes) ChatMembers(chatMembers ChatMembers) *Scenes {
	scenes.ChatMembers_ = chatMembers
	return scenes
}

func (scenes *Scenes) OnLeftChatMembers(leftChatMembers LeftChatMembers) *Scenes {
	return scenes.LeftChatMembers(leftChatMembers)
}

func (scenes *Scenes) LeftChatMembers(leftChatMembers LeftChatMembers) *Scenes {
	scenes.LeftChatMembers_ = leftChatMembers
	return scenes
}
