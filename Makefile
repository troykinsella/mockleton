PACKAGE=github.com/troykinsella/mockleton
BINARY=mockleton
COMMAND=${PACKAGE}/cmd/${BINARY}

VERSION=0.0.1

LDFLAGS=-ldflags "-X ${PACKAGE}.AppVersion=${VERSION}"

build:
	go build ${LDFLAGS} ${COMMAND}

install:
	go install ${LDFLAGS}

test:
	ginkgo ./...

dist:
	GOOS=linux  GOARCH=amd64 go build ${LDFLAGS} -o mockleton_linux_amd64 ${COMMAND}
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o mockleton_darwin_amd64 ${COMMAND}

clean:
	rm ${BINARY} || true
	rm ${BINARY}_* || true
	rm mockleton.out || true

.PHONY: build install test dist clean
