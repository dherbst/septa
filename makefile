.PHONY: *

GOLANG = golang:1.10
GOOS = darwin

all: clean pull build septaimage

clean:
	mkdir -p bin
	rm -f bin/septabot bin/narb bin/sub || true

pull:
	docker pull ${GOLANG}

build:
	docker run -i --rm -v "$(PWD)":/usr/src/myapp -w /usr/src/myapp ${GOLANG} make build-in-container

build-in-container:
	GOPATH=/usr/src/myapp go vet septabot
	GOPATH=/usr/src/myapp go test -coverprofile=coverage.out septabot
	GOOS=${GOOS} GOPATH=/usr/src/myapp go build -o bin/septabot /usr/src/myapp/src/septabot/cmd/septabot/septabot.go
	GOOS=${GOOS} GOPATH=/usr/src/myapp go build -o bin/narb /usr/src/myapp/src/septabot/cmd/narb/narb.go
	GOOS=${GOOS} GOPATH=/usr/src/myapp go build -o bin/sub /usr/src/myapp/src/septabot/cmd/sub/sub.go


septaimage: build
	docker build -t septabot:latest .

run:
	docker run -p 8080:8080 -d septabot:latest
