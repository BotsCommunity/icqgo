package API

import "net/url"

type GetFile struct {
	FileId string
}

func (bot *Bot) GetFile(properties ...any) (file ResponseGetFile) {
	var params = make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case GetFile:
			if property.FileId != "" {
				params.Set("fileId", property.FileId)
			}

		case FileID:
			params.Set("fileId", string(property))

		case string:
			params.Set("fileId", property)
		}
	}

	err := bot.Call("files/getInfo", params, &file)
	if err != nil {
		return file
	}

	return file
}
