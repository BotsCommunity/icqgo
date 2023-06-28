package API

type ResponseInfo struct {
	Info
}

type Info struct {
	OK          bool   `json:"ok"`
	Description string `json:"description"`
}
