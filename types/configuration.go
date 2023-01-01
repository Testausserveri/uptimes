package types

import "time"

type DomainUptimeRequirements struct {
	StatusCode  int    `toml:"status_code" json:"status_code"`
	ContentType string `toml:"content_type" json:"content_type"`
}

type DomainConfiguration struct {
	Name           string                   `json:"name" toml:"name"`
	Domain         string                   `json:"domain" toml:"domain"`
	Requirements   DomainUptimeRequirements `json:"requirements" toml:"requirements"`
	UpdateInterval time.Duration            `json:"update_interval" toml:"update_interval"`
}

type StatusGroupConfiguration struct {
	Name      string `toml:"name" json:"name"`
	ServePath string `toml:"serve_path" json:"serve_path"`

	Domains []DomainConfiguration `toml:"domains"`
}
