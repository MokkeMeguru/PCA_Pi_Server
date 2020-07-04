FROM golang:1.14.4
RUN apt-get update -qq && \
    apt-get install wget
WORKDIR /go/src/app