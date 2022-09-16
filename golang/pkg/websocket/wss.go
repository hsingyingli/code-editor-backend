package websocket

import (
	"log"
	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    //CheckOrigin: func(r *http.Request) bool { return true },
    ReadBufferSize:  4096,
    WriteBufferSize: 4096,
}

type WebSocketServer struct {
  clients map[*Client]bool
  rooms map[*Room]bool
  broadcast chan []byte
  register chan *Client
  unregister chan *Client
}

func (wss *WebSocketServer) run () {
  for {
		select {
		case client := <-wss.register:
		wss.clients[client] = true
		case client := <-wss.unregister:
			if _, ok := wss.clients[client]; ok {
				delete(wss.clients, client)
				close(client.send)
			}
		case message := <-wss.broadcast:
			for client := range wss.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(wss.clients, client)
				}
			}
		}
	}
}

func NewWebSocketServer() *WebSocketServer{
  wss := &WebSocketServer{
    clients: make(map[*Client]bool),
    rooms: make(map[*Room]bool),
    broadcast: make(chan []byte),
    register: make(chan *Client),
    unregister: make(chan *Client),
  }
  go wss.run()

  return wss
}


func (wss *WebSocketServer) ServeWss(ctx *gin.Context) {

    conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
    if err != nil {
        log.Println(err)
        return
    }

    sender := ctx.Request.URL.Query().Get("userID")

    client := newClient(wss, conn, sender)
    wss.register <- client
    log.Println("register")
    go client.writePump()
	  go client.readPump()
}


func (wss *WebSocketServer) findRoomByID(roomID string) *Room {
  var foundRoom *Room
  for room := range wss.rooms {
    if (room.ID == roomID) {
      foundRoom = room;
      break;
    }
  }
  return foundRoom
}

func (wss *WebSocketServer) createRoom(roomID string, owner string)*Room {

  room := NewRoom(roomID, owner)
  wss.rooms[room] = true
  go room.run()
  return room
}
