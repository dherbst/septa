all: clean pull build septaimage

clean:
	mkdir -p bin
	rm -f bin/septabot || true

pull:
	docker pull golang:1.7

build:
	docker run -i --rm -v "$(PWD)":/usr/src/myapp -w /usr/src/myapp golang:1.7 make build-in-container

build-in-container:
	GOPATH=/usr/src/myapp go vet septabot
	GOPATH=/usr/src/myapp go test -coverprofile=coverage.out septabot
	GOPATH=/usr/src/myapp go build -o bin/septabot /usr/src/myapp/src/septabot/cmd/*.go

septaimage: build
	docker build -t septabot:latest .
