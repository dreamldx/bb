PROJECT=bb

.EXPORT_ALL_VARIABLES:
GOPATH=$(shell pwd)

all:
	go build -o out/$(PROJECT) -v src/github.com/dreamldx/bb/cmd/bb.go
