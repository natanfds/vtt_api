package models

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	Room     string `json:"room"`
}
