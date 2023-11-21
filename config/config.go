package config

import (
	"log"
	"os"

	"httpserver/argparser"

	"github.com/BurntSushi/toml"
	"github.com/jessevdk/go-flags"
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
	var config Config
	_, err := os.Stat(configfile)
	if err != nil {
		config = Config{
			Server:   Server{},
			Resource: Resource{},
			Ips:      Ips{},
		}
	} else {

		if _, err := toml.DecodeFile(configfile, &config); err != nil {
			log.Fatal(err)
		}

	}

	// Parser Args
	_, errParser := argparser.Parse()
	if errParser != nil {
		switch flagsErr := errParser.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
	if argparser.Parsed.Port != 0 {
		config.Server.Port = argparser.Parsed.Port
	}
	if argparser.Parsed.Ips != "" {
		config.Ips.Read = argparser.Parsed.Ips
	}
	//if config.Resource == (Resource{}) {
	if argparser.Parsed.Resource != "" {
		config.Resource.Name = argparser.Parsed.Resource
	}
	//}

	return config
}
