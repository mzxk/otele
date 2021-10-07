package otele

import (
	"log"
	"strings"
	"time"

	"github.com/mzxk/ohttp"
	"github.com/mzxk/olog"
	"github.com/mzxk/oval"
)

type teleBot struct {
	url          string
	proxy        string
	db           *oval.KV
	updateOffset int64

	fMessage func(*Message)
	fCommand map[string]func([]string, *Message)
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

		fMessage: func(m *Message) {},
		fCommand: map[string]func([]string, *Message){},
	}
	return t, t.testBot()
}
func (t *teleBot) testBot() error {
	t.OnCommand("/echo", func(s []string, m *Message) {
		m.Reply(strings.Join(s, "-"))
	})
	s, e := t.Do("getMe", nil)
	log.Println(s, e)
	return e
}
func (t *teleBot) Do(method string, result interface{}, params ...interface{}) (string, error) {
	ohtp := ohttp.HTTP(t.url+method, params...)
	if t.proxy != "" {
		ohtp = ohtp.Proxy(t.proxy)
	}
	resp, err := ohtp.Get()
	if err != nil {
		return "", err
	}
	if result != nil {
		err = resp.JSON(&result)
	}
	return resp.String(), err
}
func (t *teleBot) SendMessage(chatid int64, text string, replyID int64) {
	var result struct {
		Ok bool
	}
RE:
	s, e := t.Do("sendMessage", &result, "chat_id", chatid, "text", text, "reply_to_message_id", replyID)
	if e != nil {
		log.Println(e)
		time.Sleep(1 * time.Second)
		goto RE
	}
	if !result.Ok {
		olog.Warn(s)
	}
}
