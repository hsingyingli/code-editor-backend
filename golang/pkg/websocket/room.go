package websocket

type Room struct {
  ID string
  owner string
  broadcast chan []byte
  register chan *Client
  unregister chan *Client
  clients map[*Client]bool
}

func NewRoom(roomID string, owner string) *Room{
  room := &Room{
    ID: roomID,
    owner: owner,
    broadcast: make(chan []byte),
    register: make(chan *Client),
    unregister: make(chan *Client),
    clients: make(map[*Client]bool),
  }
  return room 
}


func (room *Room) run() {
  for {
		select {
		case client := <-room.register:
		  room.clients[client] = true
      room.notifyClientInRoom(client.name)
		case client := <-room.unregister:
			if _, ok := room.clients[client]; ok {
				delete(room.clients, client)
			}
		case message := <-room.broadcast:
			for client := range room.clients {
				  client.send <- message
				}
			}
		}
}


func (room *Room) notifyClientInRoom(name string) {
  message := Message {
    Action: JoinRoom,
    Message: JoinRoom,
    RoomID: room.ID,
    Sender: name,
  }
  for client := range room.clients {
    client.send <- message.encode()
  }
}
