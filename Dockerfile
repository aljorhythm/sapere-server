FROM golang:1.16.4-buster AS build

COPY . /go/src/github.com/aljorhythm/sapere-server
WORKDIR /go/src/github.com/aljorhythm/sapere-server
RUN go install /go/src/github.com/aljorhythm/sapere-server

ENTRYPOINT /go/bin/sapere-server

EXPOSE 8080