
# HTTP Server

## Getting Started

Import the libraries:
`
go get github.com/BurntSushi/toml
`

Launch:
```
go run main.go
```

## Test 

```
go test ./...
```


## Build

go build main.go

## Library used


* github.com/BurntSushi/toml  : for parse toml configuration file

## Configuration file:

```toml
[server]

port = 3300

[resource]

name = "Muletto1"
```
