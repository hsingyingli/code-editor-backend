package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/hsingyingli/code-editor/pkg/api"
	"github.com/hsingyingli/code-editor/pkg/db/sqlc"
	"github.com/hsingyingli/code-editor/pkg/websocket"
)


const (
  dbDriver = "postgres"
  dbSource = "postgresql://aaron:secret@localhost:5432/code_editor?sslmode=disable" 
)

func main() {
  conn, err := sql.Open(dbDriver, dbSource)
  if err!= nil {
    log.Fatal("cannot connect to db:", err)
  }

  store := db.NewStore(conn)

  wss := websocket.NewWebSocketServer()
  server := api.NewServer(store, wss)
  log.Fatal(server.Start(":9010"))
}
