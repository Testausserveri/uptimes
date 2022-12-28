package storage

import (
	"time"

	"github.com/Testausserveri/uptimes/configuration"
)

type DomainHistoryStatus struct {
	Date         time.Time
	Up           bool
	Error        error
	ResponseTime int64
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

func (s *Storage) AddOrUpdate(domain configuration.Domain, err error, responseTime int64) {
	var found bool
	for _, storageElement := range *s {
		if storageElement.Domain.Name == domain.Name {
			storageElement.CurrentlyUp = err == nil
			if err == nil {
				storageElement.LastResponseTime = responseTime
			}
			found = true
		}
	}
	if !found {
		domainStatus := &DomainStatus{
			Domain:      domain,
			CurrentlyUp: err == nil,
		}

		domainStatus.History = append(domainStatus.History, DomainHistoryStatus{
			Date:         time.Now(),
			Up:           err == nil,
			ResponseTime: responseTime,
			Error:        err,
		})

		if err == nil {
			domainStatus.LastResponseTime = responseTime
		}

		*s = append(*s, domainStatus)
	}
}
