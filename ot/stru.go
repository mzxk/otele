package ot

import "strings"

func ParseFrom(f From) (int64, string) {
	s := strings.Join([]string{f.UserName, f.FirstName, f.LastName, f.LanguageCode}, "|")
	return f.ID, s
}
func ParseChat(f Chat) (int64, string, string) {
	s := strings.Join([]string{f.UserName, f.FirstName, f.LastName}, "|")
	return f.ID, s, f.Type
}

type From struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"username"`
	LanguageCode string `json:"language_code"`
}
type Chat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Type      string `json:"type"`
}
type Entities struct {
	Offset int64
	Length int64
	Type   string
}
type Message struct {
	MessageID int64 `json:"message_id"`
	From      From
	Chat      Chat
	Date      int64
	Text      string
	Entities  []Entities
}
type MyChatMember struct {
	UpdateID   int64 `json:"update_id"`
	ChatMember struct {
		Chat Chat
		From From
		Date int64
	}
}
type Updates struct {
	Ok     bool
	Result []Update
}
type Update struct {
	UpdateID   int64        `json:"update_id"`
	Message    Message      `json:"message"`
	ChatMember MyChatMember `json:"my_chat_member"`
}
