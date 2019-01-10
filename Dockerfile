ARG GO_VERSION=1.11
ARG ALPINE_VERSION=3.8

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder
RUN apk add --no-cache git ca-certificates

ENV CGO_ENABLED=0 GO111MODULE=on

WORKDIR /go/src/github.com/kubesail/qotm

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . /go/src/github.com/kubesail/qotm
RUN go build -a -installsuffix cgo -ldflags '-s' -o ./bin/qotm .

FROM alpine:${ALPINE_VERSION}

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/kubesail/qotm/bin/qotm ./qotm
COPY --from=builder /go/src/github.com/kubesail/qotm/favicon.ico ./favicon.ico

RUN addgroup -S kubesail && adduser -S -G kubesail kubesail && chown -R kubesail:kubesail ./qotm
USER kubesail

EXPOSE 8080

CMD ["./qotm"]
