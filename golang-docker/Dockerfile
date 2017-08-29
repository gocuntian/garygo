FROM golang:latest

ADD . /go/src/github.com/xingcuntian/go_test/golang-docker

RUN go install github.com/xingcuntian/go_test/golang-docker

EXPOSE 18080

ENTRYPOINT /go/bin/golang-docker