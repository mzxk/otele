package otele

import (
	"fmt"
	"log"
	"strings"

	"github.com/mzxk/oval"
)

type TeleBot struct {
	url          string
	proxy        string
	db           *oval.KV
	updateOffset int64

	fMessage     func(*Message)
	fCommand     map[string]func([]string, *Message)
	fCommandNote []string
}

func New(key, proxy string) (*TeleBot, error) {
	ss := strings.Split(key, ":")
	db, err := oval.NewKV("bot" + ss[0])
	if err != nil {
		panic(err)
	}
	t := &TeleBot{
		url:   "https://api.telegram.org/bot" + key + "/",
		proxy: proxy,
		db:    db,

		fMessage:     func(m *Message) {},
		fCommand:     map[string]func([]string, *Message){},
		fCommandNote: []string{},
	}
	return t, t.testBot()
}
func (t *TeleBot) testBot() error {
	t.initDefaultCmd()
	s, e := t.Do("getMe", nil)
	log.Println(s, e)
	return e
}
func (t *TeleBot) initDefaultCmd() {
	t.OnCommand("/getid", func(s []string, m *Message) {
		m.Reply(fmt.Sprintf("UserID:%d , ChatID:%d", m.FromID, m.ChatID))
	}, "Return ChatID and UserID")
	t.OnCommand("/help", func(s []string, m *Message) {
		m.Reply(strings.Join(t.fCommandNote, "\n"))
	})
}
