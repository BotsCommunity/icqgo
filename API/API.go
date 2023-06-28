package API

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	log "github.com/sirupsen/logrus"
)

type Bot struct {
	BotInfo
	Token  string       `json:"token"`
	Client *http.Client `json:"client"`
	Host   Host         `json:"host"`
	Logger *log.Logger  `json:"logger"`
}

type BotInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Host struct {
	Protocol string `json:"protocol"`
	Domain   string `json:"domain"`
}

func Create(properties ...any) *Bot {
	bot := Bot{
		Client: http.DefaultClient,
		Host: Host{
			Protocol: "https://",
			Domain:   "api.icq.net/bot/v1",
		},
	}

	for _, property := range properties {
		switch property := property.(type) {
		case string:
			bot.Token = property
		case *log.Logger:
			bot.Logger = property
		}
	}

	bot.make()
	return &bot
}

type File struct {
	Field string
	Data  *os.File
}

// Call 🔥🌅 sends a request to the VKontakte server
// 🌌❄ The first argument is the name of the method,
// 🌺🌼🌹 the second is the URL Values,
// 🦋🌹🎲 then usually a pointer to the structure where the response will be written
// 🌍🌊🐠 May return an error
func (bot *Bot) Call(method string, values url.Values, data any, file ...File) error {
	if values == nil {
		values = make(url.Values)
	}
	values.Set("token", bot.Token)

	logger := &log.Entry{}
	if bot.Logger != nil {
		logger = bot.Logger.WithFields(log.Fields{
			"botID":  bot.ID,
			"method": method,
			"values": values,
		})
	}

	URL, err := url.Parse(bot.Host.Protocol + bot.Host.Domain + "/" + method)
	if err != nil {
		logger.Error("Call ParseURL error is ", err.Error())
		return err
	}

	URL.RawQuery = values.Encode()
	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, URL.String(), nil)
	if err != nil {
		logger.Error("Call HTTP Request error is ", err.Error())
		return err
	}

	if len(file) > 0 {
		file := file[0]
		if file.Field != "" {
			var (
				buffer    = &bytes.Buffer{}
				multipart = multipart.NewWriter(buffer)
			)

			writer, err := multipart.CreateFormFile(file.Field, file.Data.Name())
			if err != nil {
				logger.Error("Call Multipart CreateFormFile error is ", err.Error())
				return err
			}

			_, err = io.Copy(writer, file.Data)
			if err != nil {
				logger.Error("Call Copy Writer error is ", err.Error())
				return err
			}

			if err := multipart.Close(); err != nil {
				logger.Error("Call Multipart Close error is ", err.Error())
				return err
			}

			request.Header.Set("Content-Type", multipart.FormDataContentType())
			request.Body = io.NopCloser(buffer)
			request.Method = http.MethodPost
		}
	}

	response, err := bot.Client.Do(request)
	if err != nil {
		logger.Error("Call HTTP Response error is ", err.Error())
		return err
	}

	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		if err := response.Body.Close(); err != nil {
			logger.Error("Call BodyClose error is ", err.Error())
		}
		return err
	}

	if bot.Logger != nil {
		info, err := json.Marshal(data)
		if err != nil {
			logger.Error("Call MarshalJSON Log error is ", err.Error())
			return err
		}

		logger.Info("Call HTTP response is ", string(info))
	}

	return err
}

func (bot *Bot) make() {
	passport := bot.Get()
	bot.ID = passport.UserID
	bot.Name = passport.Name
}
