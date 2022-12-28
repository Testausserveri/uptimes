package core

import (
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
	timeStarted := time.Now()
	err := VerifyHost(domain.Url, domain.Requirements)
	s.Storage.AddOrUpdate(domain, err, time.Since(timeStarted).Milliseconds())
	return nil
}
