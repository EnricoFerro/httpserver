package argparser

import "github.com/jessevdk/go-flags"

var Parsed Options

type Options struct {
	Port int64 `short:"p" long:"port" description:"Connect to the port specified. If not given, the default of 3030." env:"HTTPSERVER_PORT" default:"3030"`

	Resource string `short:"r" long:"resource" description:"Resource name for the device. If omitted will be ingnored" env:"HTTPSERVER_RESOURCE"`

	Ips string `short:"i" long:"ips" description:"How to read the ips if omitted is ignored"  choice:"first" choice:"last" choice:"list" choice:"0" choice:"1" choice:"2" choice:"3" choice:"4" choice:"5" choice:"6" choice:"7" choice:"8" choice:"9" choice:"10" choice:"11" choice:"12" choice:"13" choice:"14" choice:"15" choice:"16" choice:"17" choice:"18" choice:"19" choice:"20" env:"HTTPSERVER_IPS" default:"first"`
}

var parser = flags.NewParser(&Parsed, flags.Default)

func Parse() ([]string, error) {
	return parser.Parse()
}
