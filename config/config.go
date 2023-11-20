package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

const (
	First string = "first"
	Last         = "last"
	All          = "all"
)

type Config struct {
	Server   Server
	Resource Resource
	Ips      Ips
}

type Server struct {
	Port int64
}

type Resource struct {
	Name string
}

type Ips struct {
	Read string
}

// Reads info from config file
func ReadConfig() Config {
	var configfile = "config.toml"
	_, err := os.Stat(configfile)
	if err != nil {
		config := Config{
			Server:   Server{Port: 3030},
			Resource: Resource{},
			Ips:      Ips{Read: Last},
		}
		return config
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}
