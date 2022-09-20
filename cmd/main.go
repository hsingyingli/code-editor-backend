package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/hsingyingli/code-editor/pkg/api"
	"github.com/hsingyingli/code-editor/pkg/db/sqlc"
	"github.com/hsingyingli/code-editor/pkg/util"
)

func main() {

	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config file")
	}

	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.DB_USERNAME, config.DB_PASSWORD, config.DB_URL, config.DB_PORT, config.DB_TABLE)

	conn, err := sql.Open(config.DB_DRIVER, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	log.Fatal(server.Start(config.SERVER_ADDR))
}
