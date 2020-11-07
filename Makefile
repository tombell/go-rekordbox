MODFLAGS=-mod=vendor
TESTFLAGS=-cover

all: test

test:
	go test ${MODFLAGS} ${TESTFLAGS} ./...

testv:
	go test ${MODFLAGS} ${TESTFLAGS} -v ./...

.PHONY: all test testv
