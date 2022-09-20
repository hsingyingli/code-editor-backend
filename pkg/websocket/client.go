package websocket

import (
	//"bytes"
	//"fmt"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)



// Client represents the websocket client at the server
type Client struct {
    // The actual websocket connection.
    conn  *websocket.Conn
    send  chan []byte
    wss   *WebSocketServer
    name  string 
    rooms map[*Room]bool
}

func newClient(wss *WebSocketServer, conn *websocket.Conn, name string) *Client {
    return &Client{
      wss: wss, 
      conn: conn, 
      send: make(chan []byte, 256),
      name: name,
      rooms: make(map[*Room]bool),
    }
}

func (c *Client) disconnect() {
  log.Println("disconnect")
  c.wss.unregister <- c 
  for room := range c.rooms {
    room.unregister <- c
  }
  c.conn.Close()
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer c.disconnect()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, newMessage, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
    
    c.handleNewMessasge(newMessage)
	}
}


func (c *Client) handleNewMessasge(newMessage []byte) {
  var message Message;

  if err := json.Unmarshal(newMessage, &message); err != nil {
    log.Println("Unable to unmarshal json message")
  }

  log.Println(message.Action)

  switch message.Action {
  case SendMessage:
    roomID := message.RoomID
    if room := c.wss.findRoomByID(roomID); room != nil {
      room.broadcast <- message.encode()
    }

	case JoinRoom:
		c.handleJoinRoomMessage(message)

	case LeaveRoom:
    c.handleLeaveRoomMessage(message)
	}
}

func (c *Client) handleJoinRoomMessage(message Message){
  log.Println(message)
  room := c.wss.findRoomByID(message.RoomID)

  if room == nil {
    room = c.wss.createRoom(message.RoomID, c.name)
    log.Println("Create Room")
  }
  log.Println("Get Room")

  if !c.isInRoom(room) {
    log.Println("NOT IN ROOM")
		c.rooms[room] = true
		room.register <- c
	}
}


func (c *Client) isInRoom(room *Room) bool {
  if _, ok := c.rooms[room]; ok {
    return true
  }
  return false
}

func (c *Client) handleLeaveRoomMessage(message Message) {
	room := c.wss.findRoomByID(message.RoomID)
	if room == nil {
		return
	}

	if _, ok := c.rooms[room]; ok {
		delete(c.rooms, room)
	}

	room.unregister <- c
}

// writePump pumps messages from the hub to the websocket connection.
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
      //log.Printf("Write: %v", message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}






