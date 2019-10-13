.PHONY: all clean pull get build build-in-container image

GOLANG := golang:1.13
GOOS := darwin

all: clean pull lint sec build install

clean:
	rm -rf vendor/
	mkdir -p bin
	rm -f bin/septa || true

get:
	mkdir -p vendor

pull:
	docker pull ${GOLANG}

lint:
	docker run -i --rm -v  ${PWD}:/go/src/github.com/dherbst/septa -w /go/src/github.com/dherbst/septa ${GOLANG} make lint-in-container

lint-in-container:
	go get -u golang.org/x/lint/golint
	golint github.com/dherbst/septa
	golint github.com/dherbst/septa/cmd/septa/...

sec:
	docker run -it --rm -v ${PWD}:/go/src/github.com/dherbst/septa -w /go/src/github.com/dherbst/septa ${GOLANG} make sec-in-container

sec-in-container:
	go get -u github.com/securego/gosec/cmd/gosec
	gosec .

build:
	docker run -i --rm -v "$(PWD)":/go/src/github.com/dherbst/septa -w /go/src/github.com/dherbst/septa ${GOLANG} make build-in-container

build-in-container:
	GOOS=darwin go build -o bin/septa cmd/septa/*.go

install:
	cp bin/septa ~/bin/septa

image: build
	docker build -t septa:latest .
