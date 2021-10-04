package otele

func (t *teleBot) UpdateStart() {
	//TODO WSS
	var offset int64
	go func() {
		for {
			var rlt struUpdate
			t.Do("getUpdates", &rlt, "limit", 100, "offset", offset)
		}
	}()
}
func (t *teleBot) update(ok StruUpdate) {

}

type struUpdate struct {
	Ok     bool
	Result []StruUpdate
}
type StruUpdate struct {
	Message    Message
	ChatMember ChatMember
}
