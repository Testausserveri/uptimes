package main

import (
	"flag"
	"log"
	"os"

	"github.com/Testausserveri/uptimes/api"
	"github.com/Testausserveri/uptimes/config"
	"github.com/Testausserveri/uptimes/services"
	"github.com/Testausserveri/uptimes/types"
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
	lf := flag.String("l", "", "log file")
	cf := flag.String("c", "configs", "configuration directory")
	la := flag.String("a", ":8080", "HTTP server listen address")
	flag.Parse()

	if *lf != "" {
		fp, err := os.OpenFile(*lf, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer fp.Close()
		log.SetOutput(fp)
		log.Println("using", fp.Name(), "as logfile")
	}

	configs, err := config.ParseConfigs(*cf)
	if err != nil {
		log.Fatal(err)
	}

	if len(configs) == 0 {
		log.Fatal("at least one configuration file is required")
	}

	server := api.NewServer(nil, *la, router())

	for _, cfg := range configs {
		if err := config.VerifyConfig(cfg); err != nil {
			log.Fatal(err)
		}

		sg := types.NewStatusGroup(cfg)
		server.HandleStatusGroup(api.NewAPIStatusGroup(sg, cfg.ServePath))
		services.InitStatusGroupUpdater(sg)
	}

	log.Printf("starting server at %s\n", *la)
	log.Fatal(server.Start())
}
