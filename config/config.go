package config

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed config.json
var rawConfig []byte

type Config struct {
	LogLevel string       `json:"logLevel"`
	Server   ServerConfig `json:"server"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

func (cfg Config) String() string {
	m, err := json.Marshal(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	return string(m)
}

func Load() (Config, error) {
	cfg := new(Config)
	if err := json.Unmarshal(rawConfig, cfg); err != nil {
		return Config{}, err
	}

	return *cfg, nil
}
