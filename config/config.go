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
	Last  string = "last"
	List  string = "list"
	C0    string = "0"
	C1    string = "1"
	C2    string = "2"
	C3    string = "3"
	C4    string = "4"
	C5    string = "5"
	C6    string = "6"
	C7    string = "7"
	C8    string = "8"
	C9    string = "9"
	C10   string = "10"
	C11   string = "11"
	C12   string = "12"
	C13   string = "13"
	C14   string = "14"
	C15   string = "15"
	C16   string = "16"
	C17   string = "17"
	C18   string = "18"
	C19   string = "19"
	C20   string = "20"
	C21   string = "21"
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
