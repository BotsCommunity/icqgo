package update

type User struct {
	ID      string `json:"userId,omitempty"`
	Name    string `json:"firstName,omitempty"`
	Surname string `json:"lastName,omitempty"`
}
