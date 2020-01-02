project:=$(shell basename $(shell pwd))
commit:=$(shell git rev-parse --short HEAD)
importpath:=github.com/chiefy/$(project)
ts:=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

$(GOPATH)/bin/dep:
	@curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

vendor: $(GOPATH)/bin/dep
	@dep ensure

test: vendor
	@go test ./...