package config

import (
	"fmt"
	"github.com/andReyM228/lib/database"
	"github.com/andReyM228/lib/redis"
	"github.com/andReyM228/one/chain_client"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	Config struct {
		Chain  chain_client.ClientConfig `yaml:"chain"`
		DB     database.DBConfig         `yaml:"db" validate:"required"`
		HTTP   HTTP                      `yaml:"http" validate:"required"`
		Redis  redis.Config              `yaml:"cache"`
		Rabbit Rabbit                    `yaml:"rabbit"`
		Extra  Extra                     `yaml:"extra" validate:"required"`
	}

	HTTP struct {
		Port int `yaml:"port" validate:"required"`
	}

	Rabbit struct {
		Url string `yaml:"url"`
	}

	Extra struct {
		Mnemonic string `yaml:"mnemonic"`
	}
)

func ParseConfig() (Config, error) {
	file, err := os.ReadFile("./cmd/config.yaml")
	if err != nil {
		fmt.Errorf("parseConfig: %s", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		fmt.Errorf("parseConfig: %s", err)
	}

	return cfg, nil
}
