package longpoll

import (
	"github.com/botscommunity/icqgo/API"
	"github.com/botscommunity/icqgo/scene"
)

type Options struct {
	Bot      *API.Bot      `json:"bot"`
	Scenes   *scene.Scenes `json:"scenes"`
	Time     int
	UpdateID int
}

func Create(properties ...any) *Options {
	config := Options{Time: 30}

	for _, property := range properties {
		switch property := property.(type) {
		case *API.Bot:
			config.Bot = property
		case *scene.Scenes:
			config.Scenes = property
		}
	}

	return &config
}
