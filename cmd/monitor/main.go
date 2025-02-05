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
    flag.Parse()
    configPath := flag.String("cfg", "config/config.yaml", "path to config file")
    var s server.Server
    s.Init(*configPath)
    s.Run()
}
