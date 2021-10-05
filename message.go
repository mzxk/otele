package otele

type Message struct {
	UpdateID int64 `json:"update_id"`
	Message  struct {
		MessageID int64 `json:"message_id"`
		From      StruFrom
		Chat      StruChat
		Date      int64
		Text      string
		Entities  []StruEntities
	}
}

func (t *teleBot) handleMessage(msg Message) {
	if len(msg.Message.Entities) > 0 {
		if msg.Message.Entities[0].Type == "bot_command" {

		}
	}
}
func (t *teleBot) OnMessage(f func(*Message) string) {
	t.fMessage = f
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
