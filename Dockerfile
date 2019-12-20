FROM golang:1.8

WORKDIR /go/src/app
COPY . .
ENV GOBIN /go/bin
RUN go get github.com/DATA-DOG/godog/cmd/godog
