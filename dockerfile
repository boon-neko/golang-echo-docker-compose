FROM debian:stable-slim

ENV GOLANG_VERSION 1.13.4
ENV GOLANG_SHA256SUM 692d17071736f74be04a72a06dab9cac1cd759377bd85316e52b2227604c004c

ENV GOPATH=/go
ENV GO111MODULE=on
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV CGO_LDFLAGS="-Wl,-rpath=/usr/local/lib"

ENV DEBIAN_FRONTEND noninteractive
ENV TERM linux
ENV INITRD No
ENV LANG C.UTF-8

RUN apt-get update -yqq && \
    apt-get install apt-file bash -yqq && \
    apt-get install -yqq --no-install-recommends \
        make \
        gcc \
        g++ \
        git \
        curl \
        ca-certificates \
        xz-utils && \
    apt-get clean && \
    rm -rf \
        /var/lib/apt/lists/* \
        /tmp/* \
        /var/tmp/* \
        /usr/share/man \
        /usr/share/doc \
        /usr/share/doc-base && \
    \
    \
    cd /usr/local/src && \
    curl -fsSL https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz -o golang.tar.gz && \
    echo "${GOLANG_SHA256SUM} golang.tar.gz" | sha256sum -c - && \
    tar -C /usr/local -xzf golang.tar.gz && \
    rm -rf /usr/local/src/*

VOLUME /go/src
WORKDIR /go/src

RUN ln -s /usr/local/go/src/runtime /go/src/runtime
