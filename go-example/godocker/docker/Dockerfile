FROM       docker.io/golang:alpine

RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main" > /etc/apk/repositories

RUN apk add --update curl bash && \
    rm -rf /var/cache/apk/*
