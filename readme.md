
# HTTP Server

## Getting Started

Import the libraries:

```bash
go get github.com/BurntSushi/toml
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

* github.com/BurntSushi/toml  : for parse toml configuration file

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
