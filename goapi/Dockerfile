FROM golang
ADD . /data/golang/src/github.com/xingcuntian/go_test/goapi
RUN go install github.com/xingcuntian/go_test/goapi
ENTRYPOINT /data/golang/bin/goapi
EXPOSE 4000
