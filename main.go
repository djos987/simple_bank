package main

import (
	"database/sql"
	"log"

	"github.com/djos987/simple_bank/api"
	db "github.com/djos987/simple_bank/db/sqlc"
	"github.com/djos987/simple_bank/util"
	_ "github.com/lib/pq"
)

var conn *sql.DB

func main() {
	var err error

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
