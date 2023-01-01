package types

import "time"

type DomainUptimeRequirements struct {
	StatusCode  int    `toml:"status_code" json:"statusCode"`
	ContentType string `toml:"content_type" json:"contentType"`
}

type DomainConfiguration struct {
	Name           string                   `toml:"name" json:"name"`
	Domain         string                   `toml:"domain" json:"domain"`
	Requirements   DomainUptimeRequirements `toml:"requirements" json:"requirements"`
	UpdateInterval time.Duration            `toml:"update_interval" json:"updateInterval"`
}

type StatusGroupConfiguration struct {
	Name      string `toml:"name" json:"name"`
	ServePath string `toml:"serve_path" json:"servePath"`

	Domains []DomainConfiguration `toml:"domains"`
}
