package API

func (bot *Bot) Get() (get ResponseGet) {
	err := bot.Call("self/get", nil, &get)
	if err != nil {
		return get
	}

	return get
}
