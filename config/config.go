package config

import (
	"io/fs"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/Testausserveri/testausuptime/types"
)

type StatusGroup struct {
	Domains []types.DomainConfiguration
}

func ParseConfigs(p string) ([]StatusGroup, error) {
	var statusgroups []StatusGroup
	err := filepath.WalkDir(p, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		var sg StatusGroup
		if _, err := toml.DecodeFile(path, sg); err != nil {
			return err
		}

		statusgroups = append(statusgroups, sg)
		return nil
	})
	return statusgroups, err
}
