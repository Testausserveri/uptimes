package configuration

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type HTMLRequirements struct {
	ClassNames []string
	ElementIds []string
	Phrases    []string
}

type JSONRequirement struct {
	Key   string
	Value any
}

type DomainRequirements struct {
	ContentType string
	Status      int
	HtmlBody    HTMLRequirements
	JsonBody    []JSONRequirement
}

type Domain struct {
	Name         string
	Url          string
	Interval     time.Duration
	Requirements DomainRequirements
}

type Config struct {
	Domains      []Domain
	ServePath    string
	TemplateName string
}

func From(configPath string) (configs []Config) {
	err := filepath.WalkDir(configPath, func(path string, d fs.DirEntry, err error) error {
		base := filepath.Base(path)
		if strings.HasSuffix(base, ".toml") {
			var config Config
			if _, err := toml.DecodeFile(path, &config); err != nil {
				return err
			}
			configs = append(configs, config)
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	return configs
}
