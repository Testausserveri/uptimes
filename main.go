package main

import (
	"flag"
	"log"

	"github.com/Testausserveri/uptimes/api"
	"github.com/Testausserveri/uptimes/config"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
)

func router() *echo.Echo {
	r := echo.New()
	r.HideBanner = true
	r.HidePort = true
	r.Debug = false
	return r
}

func main() {
	cf := flag.String("c", "configs", "configuration directory")
	la := flag.String("a", ":8080", "HTTP server listen address")
	flag.Parse()

	configs, err := config.ParseConfigs(*cf)
	if err != nil {
		log.Fatal(err)
	}

	if len(configs) == 0 {
		log.Fatal("at least one configuration file is required")
		return
	}

	for _, cfg := range configs {
		if err := config.VerifyConfig(cfg); err != nil {
			log.Fatal(err)
		}
	}

	server := api.NewServer(&pgx.Conn{}, *la, router())
	log.Printf("starting server at %s\n", *la)
	log.Fatal(server.Start())
}
