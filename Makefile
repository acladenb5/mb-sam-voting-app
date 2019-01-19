BINARY=main

VERSION=`cat VERSION.txt`
BUILD=`git rev-parse head`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.DEFAULT_GOAL: ${BINARY}

${BINARY}:
	GOOS=linux go build -v ${LDFLAGS} -o ${BINARY} .

get:
	go get -v .

install:
	GOOS=linux go install -v ${LDFLAGS} -o ${BINARY} .

clean:
	if [ -f ${BINARY} ] ; then rm -f ${BINARY} ; fi

.PHONY: clean install

