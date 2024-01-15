.DEFAULT_GOAL := help
.PHONY: help

APP_FILE_NAME := kbot

APP=$(shell basename -s .git $(shell git remote get-url origin))

# REGISTRY=europe-docker.pkg.dev/prometheus-407701/prometheus-kbot
REGISTRY=dkzippa

BUILD_INFO := $(shell git rev-parse --short HEAD)
VERSION := $(shell git describe --tags --abbrev=0 | head -n 1)-${BUILD_INFO}

TARGET_OS := linux

GO_ARCH := $(shell uname -m)
ifeq (${GO_ARCH},aarch64)
	GO_ARCH = arm64
endif
ifeq (${GO_ARCH},x86_64)
	GO_ARCH = amd64
endif


DOCKER_IMG_NAME = ${REGISTRY}/${APP}:${VERSION}-${GO_ARCH}

D = \033[0m
R = \033[1;31m
Y = \033[1;93m


help: 		
	@printf "$YCompile Kbot ${VERSION} for ${TARGET_OS}/${GO_ARCH} $D\n"
	@echo "Usage: make [target]\n "
	@grep -B1 -E "^[a-zA-Z0-9_-]+\:([^\=]|$$)" Makefile \
	| grep -v -- -- \
	| awk 'NF' \
	| awk -F  ':' '{print $1}'
	@echo "\n"
	
	
get: 
	@echo "Getting dependencies...\n"
	go get	

build: get format 
	@printf "$RCompiling Kbot ${VERSION} for ${TARGET_OS}/${GO_ARCH}... $D\n"	
	CGO_ENABLED=0 GOOS=${TARGET_OS} GOARCH=${GO_ARCH}  go build -x -o ${APP_FILE_NAME} --ldflags="-X 'github.com/dkzippa/prometheus-kbot/cmd.appVersion=${VERSION}'"
	@echo "\nCompiled Kbot ${VERSION}, check ./${APP_FILE_NAME} \n"

lint: 
	@echo "Linting the code..."
	golangci-lint run
	@echo "\n"
	
format: 
	@echo "Formatting the code...\n"
	@gofmt -s -w ./

test:
	@echo "Running tests..."
	go test -v
	@echo "\n"

clean: 
	rm -f ./${APP_FILE_NAME}
	docker rmi ${DOCKER_IMG_NAME} --force


linux:	
	$(MAKE) build TARGET_OS=linux

macos:
	$(MAKE) build TARGET_OS=darwin

darwin: macos	

windows:
	$(MAKE) build TARGET_OS=windows

arm:
	$(MAKE) build TARGET_OS=linux TARGET_ARCH=arm64

image:	
	@echo "docker build -t ${DOCKER_IMG_NAME} --build-arg TARGET_OS=${TARGET_OS} ."
	docker build -t ${DOCKER_IMG_NAME} --build-arg TARGET_OS=${TARGET_OS} .


push:
	docker push ${DOCKER_IMG_NAME}

run:
	docker run -ti ${DOCKER_IMG_NAME} 
	
test-run: image run
