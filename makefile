.PHONY: all clean pull get build build-in-container image test test-in-container install install-local

GOLANG := golang:1.19
GOOS := darwin

all: clean pull test build

clean:
	mkdir -p bin
	rm -f bin/septa || true

pull:
	docker pull ${GOLANG}

test:
	docker run -it --rm -v ${PWD}:/go/src/github.com/dherbst/septa -w /go/src/github.com/dherbst/septa ${GOLANG} make test-in-container

test-in-container:
	go test -coverprofile=coverage.out github.com/dherbst/septa
	go tool cover -html=coverage.out -o coverage.html

build:
	docker run -i --rm -v "$(PWD)":/go/src/github.com/dherbst/septa -w /go/src/github.com/dherbst/septa ${GOLANG} make build-in-container

build-in-container:
	GOOS=darwin go build -o bin/septa cmd/septa/*.go

install-local:
	mkdir -p ~/bin
	cp bin/septa ~/bin/septa

install:
	go install github.com/dherbst/septa/cmd/septa

image: build
	docker build -t septa:latest .

# gh-release creates a new release in github and uploads the built binary.
gh-release:
	gh release create ${VERSION} 'bin/septa.zip'
