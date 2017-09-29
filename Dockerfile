FROM golang:1.9-alpine

# Package dependencies
RUN apk add --update --no-cache git libc-dev gcc

# install dependency tool
RUN go get github.com/golang/dep && go install github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/AMFDPMTE/list-api
COPY . .

# Get dependencies
RUN dep ensure -v
# Compile code
RUN go build

RUN ["./list-api"]
