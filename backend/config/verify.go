package config

import (
	"errors"
	"time"

	"github.com/Testausserveri/uptimes/types"
)

func VerifyConfig(config types.StatusGroupConfiguration) error {
	if config.Name == "" {
		return errors.New("name should not be empty")
	}

	for _, d := range config.Domains {
		if d.Name == "" {
			return errors.New("domain's name should not be empty")
		}
		if d.Domain == "" {
			return errors.New("domain should not be empty")
		}
		if d.UpdateInterval < time.Second {
			return errors.New("update interval should be at least 1 second")
		}
	}

	return nil
}
