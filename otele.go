package otele

import (
	"log"
	"strings"

	"github.com/mzxk/oval"
)

type teleBot struct {
	url          string
	proxy        string
	db           *oval.KV
	updateOffset int64

	fMessage     func(*Message)
	fCommand     map[string]func([]string, *Message)
	fCommandNote []string
}

func New(key, proxy string) (*teleBot, error) {
	ss := strings.Split(key, ":")
	db, err := oval.NewKV("bot" + ss[0])
	if err != nil {
		panic(err)
	}
	t := &teleBot{
		url:   "https://api.telegram.org/bot" + key + "/",
		proxy: proxy,
		db:    db,

		fMessage:     func(m *Message) {},
		fCommand:     map[string]func([]string, *Message){},
		fCommandNote: []string{},
	}
	return t, t.testBot()
}
func (t *teleBot) testBot() error {
	t.initDefaultCmd()
	s, e := t.Do("getMe", nil)
	log.Println(s, e)
	return e
}
func (t *teleBot) initDefaultCmd() {
	t.OnCommand("/echo", func(s []string, m *Message) {
		m.Reply(strings.Join(s, "-"))
	}, "Test Command . Just Return your message")
	t.OnCommand("/?", func(s []string, m *Message) {
		m.Reply(strings.Join(t.fCommandNote, "\n"))
	}, "This Command!")
}
