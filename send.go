package otele

import (
	"log"
	"time"

	"github.com/mzxk/ohttp"
	"github.com/mzxk/olog"
)

func (t *TeleBot) Do(method string, result interface{}, params ...interface{}) (string, error) {
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
func (t *TeleBot) SendMessage(chatid int64, text string, replyID int64) {
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

//SendFile 暂时只能发送文件,之后可能修改成发送其他类型的东西
func (t *TeleBot) SendFile(path string, id int64) error {
	return t.sendFile("sendDocument", path, "document", id)
}

//SendFile 暂时只能发送文件,之后可能修改成发送其他类型的东西
func (t *TeleBot) sendFile(method, path, field string, id int64) error {

	ohtp := ohttp.HTTP(t.url+method, "chat_id", id)
	if t.proxy != "" {
		ohtp = ohtp.Proxy(t.proxy)
	}
	resp, err := ohtp.PostFile(path, field)
	if err != nil {
		return err
	}
	var result struct {
		Ok bool
	}
	err = resp.JSON(&result)
	return err
}
