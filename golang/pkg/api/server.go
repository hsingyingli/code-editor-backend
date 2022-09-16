package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hsingyingli/code-editor/pkg/db/sqlc"
	"github.com/hsingyingli/code-editor/pkg/websocket"
)

type Server struct {
  store *db.Store
  router *gin.Engine
  wss *websocket.WebSocketServer
}


func NewServer(store *db.Store, wss *websocket.WebSocketServer) *Server {
  server := &Server{
    store: store,
    router: gin.Default(),
    wss: wss,
  }

  server.setupRouter()

  return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}


