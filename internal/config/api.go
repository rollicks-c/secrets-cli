package config

import (
	"fmt"
	"github.com/rollicks-c/configcove"
	"github.com/rollicks-c/configcove/profiles"
)

type Backends struct {
	Vault     Vault     `yaml:"vault"`
	Bitwarden Bitwarden `yaml:"bitwarden"`
}

type Vault struct {
	Address string `yaml:"address" env:"VT_ADDR"`
}

type Bitwarden struct {
	DataDir string `yaml:"dataDir" env:"BW_DATA_DIR"`
}

type Configuration struct {
	Backends Backends         `yaml:"backends"`
	Aliases  map[string]Alias `yaml:"aliases"`
	Name     string           `yaml:"name"`
}

type Alias struct {
	Command string `yaml:"command"`
	Tags    string `yaml:"tags"`
}

func Profiles() *profiles.Manager[Configuration] {
	pm := configcove.Profiles[Configuration](AppName)
	defaultConf, err := loadEmbeddedConfig()
	if err != nil {
		panic(err)
	}
	profiles.NewManager[Configuration](AppName, profiles.WithDefault[Configuration](defaultConf))
	return pm
}

func LoadVaultToken(conf Configuration) (string, bool) {
	name := fmt.Sprintf("%s.vault-token", conf.Name)
	return configcove.Store(AppName).LoadSecret(name)
}

func SaveVaultToken(conf Configuration, token string) {
	name := fmt.Sprintf("%s.vault-token", conf.Name)
	configcove.Store(AppName).SaveSecret(name, token)
}
