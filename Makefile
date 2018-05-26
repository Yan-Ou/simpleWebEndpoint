APP?=webapp
PORT?=8000
RELEASE?=0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
GOOS?=linux
GOARCH?=amd64

clean:
	rm -f ${APP}

build: clean
		CGO_ENABLED=0 GOOS=${GOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X main.Release=${RELEASE} \
		-X main.Commit=${COMMIT} -X main.BuildTime=${BUILD_TIME} \
		-X main.ProjectName=${APP}" \
		-o ${APP}

run: build
			PORT=${PORT} ./${APP}

test: run
	go test -v -race ./...
