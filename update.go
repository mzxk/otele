package otele

import (
	"fmt"
	"log"

	"github.com/mzxk/olog"
)

func (t *teleBot) UpdateStart() {
	//TODO WSS
	t.updateOffset = t.db.GetInt("offset")
	go func() {
		for {
			var rlt struUpdate
			rltStr, err := t.Do("getUpdates", &rlt, "limit", 100, "offset", t.updateOffset)
			log.Println(rltStr)
			if err != nil {
				olog.Warn(err)
			}
			if rlt.Ok {
				for i := range rlt.Result {
					go t.update(rlt.Result[i])
				}
			} else {
				olog.Warn("Telegram Result Not OK:")
			}
		}
	}()
}

func (t *teleBot) setOffset(up int64) {
	if t.updateOffset <= up {
		t.updateOffset = up + 1
		t.setOffset(t.updateOffset)
	}
	t.db.Set("offset", t.updateOffset)
}
func (t *teleBot) update(up StruUpdate) {
	if up.Message.UpdateID != 0 {
		t.setOffset(up.Message.Message.MessageID)
		t.handleMessage(up.Message)
	} else {
		fmt.Println("not handle")
	}
}

type struUpdate struct {
	Ok     bool
	Result []StruUpdate
}
type StruUpdate struct {
	UpdateID   int64      `json:"update_id"`
	Message    Message    `json:"message"`
	ChatMember ChatMember `json:"my_chat_member"`
}
