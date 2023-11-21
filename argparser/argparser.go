package argparser

import "github.com/jessevdk/go-flags"

var Parsed Options

type Options struct {
	Port int64 `short:"p" long:"port" description:"Connect to the port specified. If not given, the default of 3030." env:"HTTPSERVER_PORT" default:"3030"`

	Resource string `short:"r" long:"resource" description:"Resource name for the device. If omitted will be ingnored" env:"HTTPSERVER_RESOURCE"`

	Ips string `short:"i" long:"ips" description:"How to read the ips if omitted is ignored"  choice:"first" choice:"last" choice:"all" env:"HTTPSERVER_IPS" default:"last"`
}

var parser = flags.NewParser(&Parsed, flags.Default)

func Parse() ([]string, error) {
	return parser.Parse()
}
