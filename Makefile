project:=$(shell basename $(shell pwd))
commit:=$(shell git rev-parse --short HEAD)
importpath:=github.com/chiefy/$(project)
ts:=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

test:
	@go test ./...