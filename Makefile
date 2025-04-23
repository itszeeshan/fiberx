VERSION := 1.0.0-beta
COMMIT := $(shell git rev-parse --short HEAD)
DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

build:
	@go build -ldflags "\
		-X 'github.com/itszeeshan/fiberx/cmd.Version=$(VERSION)' \
		-X 'github.com/itszeeshan/fiberx/cmd.CommitHash=$(COMMIT)' \
		-X 'github.com/itszeeshan/fiberx/cmd.BuildDate=$(DATE)'" \
		-o bin/fiberx ./main.go
