package otele

import (
	"strings"

	"github.com/mzxk/ohttp"
	"github.com/mzxk/oval"
)

type teleBot struct {
	url          string
	proxy        string
	db           *oval.KV
	updateOffset int64

	fMessage func(*Message) string
}

func New(key, proxy string) *teleBot {
	ss := strings.Split(key, ":")
	db, err := oval.NewKV("bot" + ss[0])
	if err != nil {
		panic(err)
	}
	t := &teleBot{
		url:   "https://api.telegram.org/bot" + key + "/",
		proxy: proxy,
		db:    db,
	}
	t.UpdateStart()
	return t
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

type StruFrom struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"username"`
	LanguageCode string `json:"language_code"`
}
type StruChat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Type      string `json:"type"`
}
type StruEntities struct {
	Offset int64
	Length int64
	Type   string
}
