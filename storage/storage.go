package storage

import (
	"time"

	"github.com/Testausserveri/uptimes/configuration"
)

type DomainHistoryStatus struct {
	Date time.Time
	Up   bool
}

type DomainStatus struct {
	Domain           configuration.Domain
	CurrentlyUp      bool
	LastResponseTime int64
	History          []DomainHistoryStatus
}

type Storage []*DomainStatus

func NewStorage() Storage {
	return []*DomainStatus{}
}

func (s *Storage) AddOrUpdate(domain configuration.Domain, up bool, responseTime int64) {
	var found bool
	for _, storageElement := range *s {
		if storageElement.Domain.Name == domain.Name {
			storageElement.CurrentlyUp = up
			if up {
				storageElement.LastResponseTime = responseTime
			}
			found = true
		}
	}
	if !found {
		domainStatus := &DomainStatus{
			Domain:      domain,
			CurrentlyUp: up,
		}

		if up {
			domainStatus.LastResponseTime = responseTime
		}

		*s = append(*s, domainStatus)
	}
}
