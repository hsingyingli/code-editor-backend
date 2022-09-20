package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/hsingyingli/code-editor/pkg/db/sqlc"
	"github.com/hsingyingli/code-editor/pkg/token"
	"github.com/hsingyingli/code-editor/pkg/util"
	"github.com/hsingyingli/code-editor/pkg/websocket"
)

type Server struct {
	config     util.Config
	store      *db.Store
	router     *gin.Engine
	wss        *websocket.WebSocketServer
	tokenMaker token.Maker
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {

	tokenMaker, err := token.NewJWTMaker(config.TOKEN_STMMETRIC_KEY)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker")
	}

	wss := websocket.NewWebSocketServer()

	server := &Server{
		config:     config,
		store:      store,
		wss:        wss,
		tokenMaker: tokenMaker,
		router:     gin.Default(),
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
