package longpoll

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/botscommunity/icqgo/update"
)

func (config *Options) Listen() {
	for {
		request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf("%s%s/events/get?token=%s&pollTime=%d&lastEventId=%d", config.Bot.Host.Protocol, config.Bot.Host.Domain, config.Bot.Token, config.Time, config.UpdateID), nil)
		if err != nil {
			fmt.Println(err)
		}

		response, err := config.Bot.Client.Do(request)
		if err != nil {
			fmt.Println(err)
		}

		updates := update.Updates{}
		if err := json.NewDecoder(response.Body).Decode(&updates); err != nil {
			if err := response.Body.Close(); err != nil {
				config.Bot.Logger.Error("LongPoll Session BodyClose error is ", err.Error())
			}
			fmt.Println(err)
		}

		for _, update := range updates.Updates {
			config.Scenes.Use(config.Bot, update)
		}

		if updates.Updates != nil {
			config.UpdateID = updates.Updates[len(updates.Updates)-1].ID
		}
	}
}
