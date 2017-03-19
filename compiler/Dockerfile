FROM golang:1.8.0

MAINTAINER "Brandon Cole"

ENV PROTOCOL_BUFFERS_VERSION=3.2.0

RUN set -x && \
    apt-get update && \
    apt-get install -y --no-install-recommends \
    unzip=6.0-16+deb8u2 && \
    apt-get clean

RUN set -x && \
    mkdir -p /tmp/protoc && \
    curl -sSL https://github.com/google/protobuf/releases/download/v${PROTOCOL_BUFFERS_VERSION}/protoc-${PROTOCOL_BUFFERS_VERSION}-linux-x86_64.zip > /tmp/protoc/protoc.zip && \
    cd /tmp/protoc && \
    unzip protoc.zip -d /usr/local/lib/protobuf && \
    ln -s /usr/local/lib/protobuf/bin/protoc /usr/local/bin/protoc && \
    chmod go+rx /usr/local/bin/protoc && \
    rm -rf /tmp/protoc

# Get the source from GitHub
RUN set -x && \
    go get -u google.golang.org/grpc && \
    go get -u github.com/golang/protobuf/protoc-gen-go

WORKDIR /artisan