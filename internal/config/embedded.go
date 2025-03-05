package config

import (
	"embed"
	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var resources embed.FS

func loadEmbeddedConfig() (Configuration, error) {

	raw, err := resources.ReadFile("config.yaml")
	if err != nil {
		return Configuration{}, err
	}

	var embeddedConf Configuration
	if err := yaml.Unmarshal(raw, &embeddedConf); err != nil {
		return Configuration{}, err
	}

	return embeddedConf, nil
}
