MODFLAGS=-mod=vendor
TESTFLAGS=-cover

CGO_LDFLAGS:=-L/usr/local/opt/openssl/lib
CGO_CPPFLAGS:=-I/usr/local/opt/openssl/include

all: test

test:
	CGO_LDFLAGS=${CGO_LDFLAGS} CGO_CPPFLAGS=${CGO_CPPFLAGS} go test ${MODFLAGS} ${TESTFLAGS} ./...

testv:
	CGO_LDFLAGS=${CGO_LDFLAGS} CGO_CPPFLAGS=${CGO_CPPFLAGS} go test ${MODFLAGS} ${TESTFLAGS} -v ./...

.PHONY: all test testv
