package types

import "time"

type DomainUptimeRequirements struct {
	StatusCode  int    `toml:"status_code"`
	ContentType string `toml:"content_type"`
}

type DomainConfiguration struct {
	UpdateInterval time.Duration            `toml:"update_interval"`
	Name           string                   `toml:"name"`
	Domain         string                   `toml:"domain"`
	Requirements   DomainUptimeRequirements `toml:"requirements"`
}

type StatusGroupConfiguration struct {
	Name      string `toml:"name"`
	ServePath string `toml:"serve_path"`

	Domains []DomainConfiguration `toml:"domains"`
}
