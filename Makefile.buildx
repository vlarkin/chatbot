GITHUB=github.com/vlarkin
APP=chatbot
REGISTRY=europe-docker.pkg.dev/skillful-fx-417519/docker-images
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
IMAGE=${REGISTRY}/${APP}:${VERSION}

go_version=$(word 4, $(shell go version))

ifndef TARGETOS
	TARGETOS=$(word 1,$(subst /, , $(go_version)))
endif

ifndef TARGETARCH
	TARGETARCH=$(word 2,$(subst /, , $(go_version)))
endif

get:
	go get

format:
	gofmt -s -w ./

lint:
	golangci-lint run

test:
	go test -v

linux: format get
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ${APP} -ldflags "-X="${GITHUB}/${APP}/cmd.appVersion=${VERSION}

windows: format get
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -o ${APP} -ldflags "-X="${GITHUB}/${APP}/cmd.appVersion=${VERSION}

macos: format get
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -v -o ${APP} -ldflags "-X="${GITHUB}/${APP}/cmd.appVersion=${VERSION}

arm: format get
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -v -o ${APP} -ldflags "-X="${GITHUB}/${APP}/cmd.appVersion=${VERSION}

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o ${APP} -ldflags "-X="${GITHUB}/${APP}/cmd.appVersion=${VERSION}

image:
	docker buildx build --platform linux/amd64,linux/arm64 -t ${IMAGE} --push .

clean:
	rm -rf chatbot
	docker rmi -f ${IMAGE}
