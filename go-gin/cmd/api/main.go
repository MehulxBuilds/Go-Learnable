// config -> db -> router -> run server

package main

import (
	"fmt"
	"log"
	"go-gin/internal/config"
	"go-gin/internal/db"
	"go-gin/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error")
	}

	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("db error")
	}

	defer func() {

		if err := db.Disconnect(client); err != nil {
			log.Printf("mongo disconnect error: %v", err)
		}

	}()

	router := server.NewRouter(database)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(addr); err != nil {
		log.Fatalf("server failed")
	}
	
}
