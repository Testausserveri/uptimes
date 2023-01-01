package config

import (
	"errors"

	"github.com/Testausserveri/uptimes/types"
)

func VerifyConfig(config types.StatusGroupConfiguration) error {
	if config.Name == "" {
		return errors.New("name should not be empty")
	}

	return nil
}
