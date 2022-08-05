# Constants

PROJECT_NAME = 'website-checker-backend'

ifeq ($(OS),Windows_NT) 
    DETECTED_OS := Windows
else
    DETECTED_OS := $(shell sh -c 'uname 2>/dev/null || echo Unknown')
endif

# Help

.SILENT: help
help:
	@echo
	@echo "Usage: make [command]"
	@echo
	@echo "Commands:"
	@echo " rename-project name={name}    	Rename project"	
	@echo	
	@echo " build                    		Build server"
	@echo " run                    			Run server"
	@echo
	@echo " docker-up                     	Up docker services"
	@echo " docker-down                   	Down docker services"
	@echo
	@echo " fmt                           	Format source code"
	@echo " test                          	Run unit tests"
	@echo

# Build

.SILENT: rename-project
rename-project:
    ifeq ($(name),)
		@echo 'new project name not set'
    else
        ifeq ($(DETECTED_OS),Darwin)
			@grep -RiIl '$(PROJECT_NAME)' | xargs sed -i '' 's/$(PROJECT_NAME)/$(name)/g'
        endif

        ifeq ($(DETECTED_OS),Linux)
			@grep -RiIl '$(PROJECT_NAME)' | xargs sed -i 's/$(PROJECT_NAME)/$(name)/g'
        endif

        ifeq ($(DETECTED_OS),Windows)
			@grep 'target is not implemented on Windows platform'
        endif
    endif

.SILENT: build
build:
	@go build -o ./bin/server ./main.go
	@echo executable file \"server\" saved in ./bin/server

# Test

.SILENT: lint
lint:
	@golangci-lint run

.SILENT: test
test:
	go test -cover ./services ./controllers -coverprofile=coverage.out -v -test.v

# Docker

.SILENT: docker-up
docker-up:
	@docker compose up -d

.SILENT: docker-down
docker-down:
	@docker compose down

# Format

.SILENT: fmt
fmt:
	@go fmt ./...

# Run server
.SILENT: run
run:
	@go build -o ./bin/server ./main.go
	@./bin/server server

.SILENT: cover
cover:
	go tool cover -html=./coverage.out -o cover.html && open cover.html

# Inject dependencies
.SILENT: wire
wire:
	go generate ./injection

# Create mocks
.SILENT: lzm
lzm:
	go generate ./repositories ./services


# Default

.DEFAULT_GOAL := help