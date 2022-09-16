package websocket
import (
  "encoding/json"
  "log"
)

type Message struct {
  Action string `json:"action"`
  Message string `json:"message"`
  RoomID string `json:"roomID"`
  Sender string `json:"sender"`
}


func (msg *Message) encode() []byte {
  json, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	return json
}
