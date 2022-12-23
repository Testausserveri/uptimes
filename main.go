package main

import (
	"flag"
	"log"
	"time"

	"github.com/Testausserveri/uptimes/configuration"
	"github.com/Testausserveri/uptimes/engine"
	"github.com/Testausserveri/uptimes/front"
	"github.com/romeq/jobscheduler"
)

func main() {
	configFile := flag.String("c", "configs", "configuration directory")
	flag.Parse()

	configurations := configuration.From(*configFile)

	for _, configFile := range configurations {
		statusGroup := engine.NewGroup(configFile)

		var jobs []jobscheduler.Job
		for _, domain := range statusGroup.Config.Domains {
			worker := newfetchworker(domain.Interval, fetch(statusGroup, domain))
			jobs = append(jobs, worker)
		}

		front.InitRoute(statusGroup)
		go jobscheduler.Run(jobs)
	}

	if err := front.Serve("localhost", 8080); err != nil {
		log.Fatalln(err)
	}
}

func newfetchworker(duration time.Duration, fn func()) jobscheduler.Job {
	return jobscheduler.NewJob(0, duration, fn, true)
}

func fetch(group engine.StatusGroup, domain configuration.Domain) func() {
	return func() {
		if err := group.UpdateDomain(domain); err != nil {
			log.Println(err)
		}
	}
}
