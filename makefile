.PHONY: all clean pull get build build-in-container image test test-in-container sec sec-in-container lint lint-in-container

GOLANG := golang:1.16
GOOS := darwin

VERSION := 1.1.3
GIT_HASH = $(shell git rev-parse --short HEAD)
LDFLAGS := "-X github.com/dherbst/septa.GitHash=${GIT_HASH} -X github.com/dherbst/septa.Version=${VERSION}"

all: clean pull lint sec test build install

clean:
	mkdir -p bin
	rm -f bin/septa || true

pull:
	docker pull ${GOLANG}

lint:
	docker run -i --rm -v ${PWD}:/go/src/github.com/dherbst/septa -w /go/src/github.com/dherbst/septa ${GOLANG} make lint-in-container

lint-in-container:
	go get -u golang.org/x/lint/golint
	golint github.com/dherbst/septa
	golint github.com/dherbst/septa/cmd/septa/...

sec:
	docker run -it --rm -v ${PWD}:/go/src/github.com/dherbst/septa -w /go/src/github.com/dherbst/septa ${GOLANG} make sec-in-container

sec-in-container:
	go get -u github.com/securego/gosec/cmd/gosec
	gosec .

test:
	docker run -it --rm -v ${PWD}:/go/src/github.com/dherbst/septa -w /go/src/github.com/dherbst/septa ${GOLANG} make test-in-container

test-in-container:
	go test -ldflags ${LDFLAGS} -coverprofile=coverage.out github.com/dherbst/septa
	go tool cover -html=coverage.out -o coverage.html

build:
	docker run -i --rm -v "$(PWD)":/go/src/github.com/dherbst/septa -w /go/src/github.com/dherbst/septa ${GOLANG} make build-in-container

build-in-container:
	GOOS=darwin go build -o bin/septa -ldflags ${LDFLAGS} cmd/septa/*.go

install:
	mkdir -p ~/bin
	cp bin/septa ~/bin/septa

install-local:
	go install -ldflags ${LDFLAGS} github.com/dherbst/septa/cmd/septa

image: build
	docker build -t septa:latest .

# gh-release creates a new release in github and uploads the built binary.
gh-release:
	gh release create ${VERSION} 'bin/septa.zip'
