package main

import (
	"dockMon/internal/server"
	"flag"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	configPath := flag.String("c", "config/config.yaml", "path to config file")
	flag.Parse()
	var s server.Server
	s.Init(*configPath)
	s.Run()
}
