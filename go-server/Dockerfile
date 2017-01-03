FROM alpine:3.4
MAINTAINER Byron Mansfield <byron@byronmansfield.com>

RUN apk add --no-cache ca-certificates

ENV GOLANG_VERSION=1.7.1
ENV GOLANG_SRC_URL=https://golang.org/dl/go${GOLANG_VERSION}.src.tar.gz \
    GOLANG_SRC_SHA256=2b843f133b81b7995f26d0cb64bbdbb9d0704b90c44df45f844d28881ad442d3

RUN set -ex && \
    apk add --no-cache --virtual .build-deps \
    bash \
    gcc \
    musl-dev \
    openssl \
    curl \
    git \
    go && \
    export GOROOT_BOOTSTRAP="$(go env GOROOT)" && \
    wget -q ${GOLANG_SRC_URL} -O golang.tar.gz && \
    echo "${GOLANG_SRC_SHA256}  golang.tar.gz" | sha256sum -c - && \
    tar -C /usr/local -xzf golang.tar.gz && \
    rm golang.tar.gz && \
    cd /usr/local/go/src && \
    ./make.bash

ENV GOPATH=/go
ENV PATH=${GOPATH}/bin:/usr/local/go/bin:${PATH}

RUN mkdir -p "${GOPATH}/src" "${GOPATH}/bin" "${GOPATH}/src/github.com/byronmansfield/go-server" && \
    chmod -R 777 "${GOPATH}"
WORKDIR ${GOPATH}/src/github.com/byronmansfield/go-server

COPY . ${WORKDIR}

RUN go get && go install

RUN apk del .build-deps

EXPOSE 8080
ENTRYPOINT [ "/go/bin/go-server" ]
