FROM circleci/golang:1.10.3

MAINTAINER Jumpei Matsuda <jmatsu.drm@gmail.com>

RUN mkdir -p $GOPATH/src/github.com/jmatsu/dpg
WORKDIR $GOPATH/src/github.com/jmatsu/dpg

ADD . $GOPATH/src/github.com/jmatsu/dpg

RUN go get -v -t -d ./...
RUN go install .
