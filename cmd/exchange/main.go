package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/RezaMokaram/ExchangeService/config"
	"github.com/RezaMokaram/ExchangeService/database"
	"github.com/RezaMokaram/ExchangeService/server"
)

func main() {
	var path string
	flag.StringVar(&path, "cleanenv", "./config/config.json", "path to clean env config file")
	flag.Parse()

	cfg, err := config.LoadConfig(path)
	if err != nil {
		log.Fatalf("load config failed: %v\n", err.Error())
	}

	db, err := database.NewConnection(&cfg.Postgres)
	if err != nil {
		log.Fatalf("db connection failed: %v\n", err.Error())
	}

	fmt.Println("Database operations done.")

	e := server.NewServer()
	server.RunServer(e, db, &cfg.App)
}