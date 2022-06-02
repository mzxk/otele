package otele

import (
	"log"
	"sync"
	"time"

	"github.com/mzxk/olog"
	"github.com/mzxk/otele/ot"
)

func (t *TeleBot) Start() {
	//TODO WSS
	t.updateOffset = t.db.GetInt("offset")
	go func() {
		for {
			var rlt ot.Updates
			rltStr, err := t.Do("getUpdates", &rlt, "limit", 100, "offset", t.updateOffset)
			if err != nil {
				olog.Warn(err)
			}
			if rlt.Ok {
				var wg sync.WaitGroup
				if len(rlt.Result) > 0 {
					log.Println(rltStr)
				}
				t.setOffset(rlt.Result)
				for i := range rlt.Result {
					wg.Add(1)
					temp := rlt.Result[i]
					go func() {
						t.update(temp)
						wg.Done()
					}()
				}
				wg.Wait()
			} else {
				olog.Warn("Telegram Result Not OK:")
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

func (t *TeleBot) setOffset(up []ot.Update) {
	if len(up) == 0 {
		return
	}
	id := up[len(up)-1].UpdateID
	t.db.Set("offset", id+1)
	t.updateOffset = id + 1
}
func (t *TeleBot) update(up ot.Update) {
	// oval.PrintStruct(up)
	if up.UpdateID != 0 {
		t.handleMessage(up.Message)
	} else {
		log.Println("not handle", up)
	}
}
