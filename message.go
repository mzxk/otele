package otele

import (
	"strings"

	"github.com/mzxk/olog"
	ot "github.com/mzxk/otele/ot"
)

type Message struct {
	MessageID int64
	Time      int64
	FromID    int64
	FromName  string
	ChatID    int64
	ChatName  string
	ChatType  string
	Text      string
	IsCommand bool

	bot *teleBot
}

func (t *Message) Reply(text string) {
	t.bot.SendMessage(t.ChatID, text, t.MessageID)
}
func (t *teleBot) newMessage(msg ot.Message) *Message {
	iscmd := false
	if len(msg.Entities) > 0 {
		if msg.Entities[0].Type == "bot_command" {
			iscmd = true
		}
	}
	fid, fname := ot.ParseFrom(msg.From)
	cid, cname, ctype := ot.ParseChat(msg.Chat)
	return &Message{
		MessageID: msg.MessageID,
		Time:      msg.Date,
		FromID:    fid,
		FromName:  fname,
		ChatID:    cid,
		ChatName:  cname,
		ChatType:  ctype,
		IsCommand: iscmd,
		Text:      msg.Text,

		bot: t,
	}
}
func (t *teleBot) handleMessage(msg ot.Message) {
	m := t.newMessage(msg)
	if m.IsCommand {
		ss := strings.Split(m.Text, " ")
		if len(ss) > 0 {
			if f, ok := t.fCommand[ss[0]]; ok {
				f(ss[1:], m)
			} else {
				olog.Warn("NotHandleCmd:", msg)
			}
		}
	} else {
		t.fMessage(m)
	}
}
func (t *teleBot) OnMessage(f func(*Message)) {
	t.fMessage = f
}
func (t *teleBot) OnCommand(cmd string, f func([]string, *Message)) {
	t.fCommand[cmd] = f
}

// "update_id": 92270932,
// "message": {
//   "message_id": 2,
//   "from": {
// 	"id": 735377945,
// 	"is_bot": false,
// 	"first_name": "NO",
// 	"last_name": "NO",
// 	"username": "q122411302019",
// 	"language_code": "zh-hans"
//   },
//   "chat": {
// 	"id": 735377945,
// 	"first_name": "NO",
// 	"last_name": "NO",
// 	"username": "q122411302019",
// 	"type": "private"
//   },
//   "date": 1633190201,
//   "text": "/start",
//   "entities": [
// 	{
// 	  "offset": 0,
// 	  "length": 6,
// 	  "type": "bot_command"
// 	}
//   ]
// }
