ARG GO_VERSION=1.14.1

FROM golang:${GO_VERSION} AS builder

ARG BUILD_MODE="install"
ARG BUILD_ARGS=''
ARG BUILD_PKG="./..."
ARG BUILD_TAG=''

ENV CGO_ENABLED=0

WORKDIR /build
ADD go.mod go.sum /build/
RUN go mod download
ADD . .
RUN go $BUILD_MODE $BUILD_ARGS $BUILD_TAG -v $BUILD_PKG

FROM alpine:3
ARG KUSTOMIZE_VERSION=3.0.3

RUN apk add --update --no-cache bash ca-certificates curl git jq openssh

COPY --from=builder /go/bin/kustomize-check-action /root/kustomize-check-action

RUN curl -Lo kustomize https://github.com/kubernetes-sigs/kustomize/releases/download/v{$KUSTOMIZE_VERSION}/kustomize_${KUSTOMIZE_VERSION}_linux_amd64 && chmod +x kustomize && mv kustomize /usr/local/bin
RUN curl -Lo kubeval https://github.com/instrumenta/kubeval/releases/latest/download/kubeval-linux-amd64.tar.gz && chmod +x kubeval && mv kubeval /usr/local/bin

WORKDIR /root/
ENTRYPOINT ["/root/kustomize-check-action"]
