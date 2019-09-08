FROM golang:1.13-stretch

LABEL maintainer="jmatsu.drm@gmail.com"

RUN mkdir -p $GOPATH/src/github.com/jmatsu/dpg

WORKDIR $GOPATH/src/github.com/jmatsu/dpg

ADD . $GOPATH/src/github.com/jmatsu/dpg

RUN go get -v -t -d ./...
RUN go install .
