FROM golang:1.8

RUN mkdir /go/src/demo-cli

WORKDIR /go/src/demo-cli

COPY . ./

ENV TEST_ENV_ITEM=dockerfile
ENV ANOTHER_ENV_ITEM="more from dockerfile"

RUN go get ./... && go build

ENTRYPOINT ["./demo-cli"]