#FROM golang:latest as build
#COPY . /usr/local/go/src/xilixili
#WORKDIR /usr/local/go/src/xilixili
#RUN  go build -o xili-server
FROM alpine

ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV GIN_MODE="release"
ENV LOG_LEVEL="release"

#SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build main.go
COPY xili-server /usr/bin/xili-server
RUN chmod +x /usr/bin/xili-server

ENTRYPOINT ["xili-server"]

