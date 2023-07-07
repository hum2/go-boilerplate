.PHONY: setup build build_release start_local lint lint_openapi generate_all generate_oapi generate_ent generate_wire

export PROJECT_DIR := $(shell pwd)

setup:
	go get -u ./...

build:
	go build -o bin/app ./cmd/server

build_release:
	go build -ldflags="-s -w" -trimpath -o bin/app ./cmd/server

start_local:
	air -c .air.toml

lint:
	npx --yes @stoplight/spectral-cli lint openapi/oapi-codegen.gen.yaml --ruleset openapi/.spectral.json
	./bin/golangci-lint run ./...

lint_openapi:
	npx --yes @stoplight/spectral-cli lint openapi/oapi-codegen.gen.yaml --ruleset openapi/.spectral.json

generate_all:
	ycl -i openapi/oapi-codegen.yamlycl -i openapi/oapi-codegen.yaml
	go generate ./... > /dev/null

generate_oapi:
	ycl -i openapi/oapi-codegen.yamlycl -i openapi/oapi-codegen.yaml
	go generate ./internal/controller/... > /dev/null

generate_ent:
	go generate ./ent/... > /dev/null

generate_wire:
	wire cmd/server/wire.go
