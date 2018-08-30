PROJECT=bb

.EXPORT_ALL_VARIABLES:
GOPATH=$(shell pwd)

all:
	go build -o $(PROJECT) -v src/github.com/dreamldx/bb/cmd/bb.go
