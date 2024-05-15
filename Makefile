
APP=chatbot
GITHUB_PROJECT=github.com/vlarkin
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)

go_version=$(word 4, $(shell go version))

ifndef TARGETOS
	TARGETOS=$(word 1,$(subst /, , $(go_version)))
endif

ifndef TARGETARCH
	TARGETARCH=$(word 2,$(subst /, , $(go_version)))
endif

ifndef DOCKER_REGISTRY
	DOCKER_REGISTRY=europe-docker.pkg.dev/skillful-fx-417519/docker-images
endif

IMAGE=${DOCKER_REGISTRY}/${APP}:${VERSION}

get:
	go get

format:
	gofmt -s -w ./

lint:
	golangci-lint run

tests:
	go test -v

linux: format get
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ${APP} -ldflags "-X="${GITHUB_PROJECT}/${APP}/cmd.appVersion=${VERSION}

windows: format get
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -o ${APP} -ldflags "-X="${GITHUB_PROJECT}/${APP}/cmd.appVersion=${VERSION}

macos: format get
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -v -o ${APP} -ldflags "-X="${GITHUB_PROJECT}/${APP}/cmd.appVersion=${VERSION}

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o ${APP} -ldflags "-X="${GITHUB_PROJECT}/${APP}/cmd.appVersion=${VERSION}

image:
	docker build --platform linux/${TARGETARCH} -t ${IMAGE}-linux-${TARGETARCH} .

push:
	docker push ${IMAGE}-linux-${TARGETARCH}

dockerize: image push

clean:
	rm -rf chatbot
