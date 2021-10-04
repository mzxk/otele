package otele

import (
	"github.com/mzxk/ohttp"
)

type teleBot struct {
	url   string
	proxy string
}

func New(key, proxy string) *teleBot {
	return &teleBot{
		url:   "https://api.telegram.org/bot" + key + "/",
		proxy: proxy,
	}
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
