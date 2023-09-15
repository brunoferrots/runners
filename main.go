package main

import (
	"log"
	"p_runner/config"
	"p_runner/server"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting pRunner App")
	log.Println("Initializing configuration")
	config := config.InitConfig("runners")

	log.Println("Initializing database")
	dbHandle := server.InitDatabase(config)

	log.Println("Initalizing server")
	httpServer := server.InitHttpServer(config, dbHandle)
	httpServer.Start()
}