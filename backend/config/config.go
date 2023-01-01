package config

import (
	"io/fs"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/Testausserveri/uptimes/types"
)

func ParseConfigs(p string) ([]types.StatusGroupConfiguration, error) {
	var statusgroups []types.StatusGroupConfiguration
	err := filepath.WalkDir(p, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if err != nil {
			return err
		}

		var sg types.StatusGroupConfiguration
		if _, err := toml.DecodeFile(path, &sg); err != nil {
			return err
		}

		statusgroups = append(statusgroups, sg)
		return nil
	})
	return statusgroups, err
}
