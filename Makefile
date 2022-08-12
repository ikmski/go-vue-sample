#===============================================================================
# Valiables
#===============================================================================

#-------------------------------------------------------------------------------
# Application
#-------------------------------------------------------------------------------
APP_NAME := app
APP_REPO := github.com/ikmski/go-vue-sample

APP_SOURCES := $(shell find . -name "*.go" -type f)

#===============================================================================
# Tasks
#===============================================================================

#-------------------------------------------------------------------------------
# Application
#-------------------------------------------------------------------------------

.PHONY: dev frontend build clean init test

default: dev

dev: frontend
	air

frontend:
	npm --prefix ./frontend run build

build: bin/$(APP_NAME)

bin/$(APP_NAME): $(APP_SOURCES) frontend
	go build \
		-a -v \
		-tags netgo \
		-installsuffix netgo \
		-ldflags "$(LDFLAGS)" \
		-o $@ \

clean:
	go clean -i ./...
	rm -rf bin/*

init:
	if [ ! -f go.mod ]; then \
		go mod init $(APP_REPO); \
	fi
	go install github.com/Songmu/make2help/cmd/make2help@latest
	go install github.com/cosmtrek/air@latest

download-deps:
	go mod download

update-deps:
	go mod tidy

test:
	go test -v $(APP_REPO)/...

