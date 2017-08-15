all: clean pull build septaimage

clean:
	mkdir -p bin
	rm -f bin/septabot bin/narb bin/sub || true

pull:
	docker pull golang:1.8

build:
	docker run -i --rm -v "$(PWD)":/usr/src/myapp -w /usr/src/myapp golang:1.7 make build-in-container

build-in-container:
	GOPATH=/usr/src/myapp go vet septabot
	GOPATH=/usr/src/myapp go test -coverprofile=coverage.out septabot
	GOPATH=/usr/src/myapp go build -o bin/septabot /usr/src/myapp/src/septabot/cmd/septabot/septabot.go
	GOPATH=/usr/src/myapp go build -o bin/narb /usr/src/myapp/src/septabot/cmd/narb/narb.go
	GOPATH=/usr/src/myapp go build -o bin/sub /usr/src/myapp/src/septabot/cmd/sub/sub.go


septaimage: build
	docker build -t septabot:latest .

run:
	docker run -p 8080:8080 -d septabot:latest
