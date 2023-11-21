
# HTTP Server

## Getting Started

Import the libraries:

```bash
go get github.com/BurntSushi/toml
go get github.com/jessevdk/go-flags
```

Launch:

```bash
go run main.go
```

## Test

```bash
go test ./...
```


## Build

go build main.go

## Library used

* [https://github.com/BurntSushi/toml](github.com/BurntSushi/toml)  : for parse toml configuration file
* [https://github.com/jessevdk/go-flags](github.com/jessevdk/go-flags) : For parse the command line

## Configuration file

```toml
[server]

port = 3300

[resource]

name = "resuorce1"


[Ips] 

# Allowed "first", "last", "all"
read = "first"
```

## 

Usage:
  httpserver [OPTIONS]

Application Options:
  -p, -port:                 Connect to the port specified. If not given, the default of 3030. (default: 3030) [$HTTPSERVER_PORT]
  -r, -resource:             Resource name for the device. If omitted will be ingnored [$HTTPSERVER_PORT]
  -i, -ips:[first|last|all]  How to read the ips if omitted is ignored (default: last) [$HTTPSERVER_PORT]

Help Options:
  -?                         Show this help message
  -h, -help                  Show this help message
