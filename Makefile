PROJECT                =permissions-service
PROJECT_DIR            =$(shell pwd)
PATH_MAIN              =${PROJECT_DIR}/cmd/app/main.go
OS                     := $(shell go env GOOS)
ARCH                   := $(shell go env GOARCH)
BUILD_FLAGS            =-ldflags \"-s\" -a -installsuffix cgo
PROJ_PORT              =3000
PROJ_ENV_VARS          =APP_PORT=${PROJ_PORT}

default: run

dependencies:
	go get github.com/golang/dep/cmd/dep
	dep ensure

run: dependencies
	$(PROJ_ENV_VARS) go run cmd/app/main.go

test: dependencies
	$(PROJ_ENV_VARS) go test -v -short ./...

integration_test: dependencies
	docker-compose build
	docker-compose up -d
	docker-compose run tester

build:
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -o main $(BUILD_FLAGS) $(PATH_MAIN)
