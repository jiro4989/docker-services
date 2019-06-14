FROM golang:1.12.5-alpine3.9 as build-stage

RUN go version && echo $GOPATH

ENV GO111MODULE off
RUN apk add --update git
ENV GO111MODULE on

WORKDIR /go/src/github.com/greymd/ojichat
# COPY Makefile /go/src/github.com/greymd/ojichat
RUN make

FROM alpine:latest as exec-stage
COPY --from=build-stage /go/src/github.com/greymd/ojichat/bin/ojichat /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/ojichat"]
