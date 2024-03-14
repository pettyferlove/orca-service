BINARY ?= orca-service
GOARCH ?= amd64
MEKE ?= make

VERSION ?= 1.0.0
BUILD ?=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.PHONY: all clean

all: $(MAKE) clean $(MAKE) darwin $(MAKE) linux $(MAKE) windows

windows: env GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}_windows_${GOARCH}.exe

darwin: env GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}_darwin_${GOARCH}

linux: env GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}_linux_${GOARCH}

clean: if ls ${BINARY}_* 1> /dev/null 2>&1; then rm ${BINARY}_* ; fi
