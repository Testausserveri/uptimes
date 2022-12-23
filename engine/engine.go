package engine

import (
	"net/http"
	"time"

	"github.com/Testausserveri/uptimes/configuration"
	"github.com/Testausserveri/uptimes/storage"
)

type StatusGroup struct {
	Storage *storage.Storage
	Config  *configuration.Config
}

func NewGroup(config configuration.Config) StatusGroup {
	storage := storage.NewStorage()
	return StatusGroup{
		Storage: &storage,
		Config:  &config,
	}
}

func (s *StatusGroup) UpdateDomain(domain configuration.Domain) error {
	var (
		alive  bool = true
		client http.Client
	)

	req, err := http.NewRequest("GET", domain.Url, nil)
	if err != nil {
		return err
	}

	timeStarted := time.Now()
	if _, err = client.Do(req); err != nil {
		alive = false
	}

	s.Storage.AddOrUpdate(domain, alive, time.Since(timeStarted).Milliseconds())
	return nil
}
