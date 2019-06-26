FROM golang:1.12.5-alpine3.9 as build-stage

ENV GO111MODULE on
RUN go version \
    && echo $GOPATH \
    && apk add --no-cache git \
    && git clone https://github.com/jiro4989/docker-services /root/go/src/github.com/jiro4989/docker-services \
    && cd /root/go/src/github.com/jiro4989/docker-services/03_proxy/app \
    && go build -o /usr/local/bin/app .

FROM alpine:latest as exec-stage
COPY --from=build-stage /usr/local/bin/app /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/app"]
