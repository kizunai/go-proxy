package lib

import (
	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Port int `toml:"port"`
}

func Config() *Configuration {
	var config Configuration

	// Load the configuration from a file.
	if _, err := toml.DecodeFile("./conf/proxy.toml", &config); err != nil {
		ProxyLogger().Error(err)
	}
	return &config
}
