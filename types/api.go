package types

import (
	"time"

	"github.com/google/uuid"
)

type StatusMetadata struct {
	Reachable bool      `json:"reachable"`
	Error     string    `json:"error"`
	Date      time.Time `json:"date"`
}

type Domain struct {
	UUID          uuid.UUID           `json:"uuid"`
	StatusGroupId uuid.UUID           `json:"statusgroupId"`
	Configuration DomainConfiguration `json:"configuration"`

	History    *[]StatusMetadata `json:"history"`
	LastStatus *StatusMetadata   `json:"lastStatus"`
}

func NewDomain(d DomainConfiguration, sgUuid uuid.UUID) *Domain {
	return &Domain{
		UUID:          uuid.New(),
		StatusGroupId: sgUuid,
		Configuration: d,
		LastStatus:    &StatusMetadata{},
		History:       &[]StatusMetadata{},
	}
}

type StatusGroup struct {
	UUID    uuid.UUID `json:"uuid"`
	Name    string    `json:"name"`
	Domains *[]Domain `json:"domains"`
}

func NewStatusGroup(s StatusGroupConfiguration) *StatusGroup {
	sgUuid := uuid.New()
	var domains []Domain
	for _, domain := range s.Domains {
		domains = append(domains, *NewDomain(domain, sgUuid))
	}

	return &StatusGroup{
		UUID:    sgUuid,
		Name:    s.Name,
		Domains: &domains,
	}
}
