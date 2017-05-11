FROM       docker.io/golang:alpine
#FROM       docker.io/alpine:latest

MAINTAINER xingcuntian <xingcuntian@qq.com>

RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main" > /etc/apk/repositories

RUN apk add --update curl bash && rm -rf /var/cache/apk/*

RUN mkdir -p /data/go

ADD http.go  /data/go/

WORKDIR /data/go/

RUN go build .

EXPOSE 8080

ENTRYPOINT ["/data/go/go"]
