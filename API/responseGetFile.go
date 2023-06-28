package API

type ResponseGetFile struct {
	Type     string `json:"type,omitempty"`
	Size     int    `json:"size,omitempty"`
	FileName string `json:"filename,omitempty"`
	URL      string `json:"url,omitempty"`
}
