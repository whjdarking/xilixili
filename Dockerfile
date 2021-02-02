FROM golang as build
COPY . /usr/local/go/src/xilixili
WORKDIR /usr/local/go/src/xilixili
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine

ENV MYSQL_DSN="root:WHJ1qaz2wsx@tcp(127.0.0.1:3306)/xilixili?charset=utf8mb4&parseTime=True&loc=Local"
ENV REDIS_ADDR="127.0.0.1:6379"
ENV REDIS_PW="1qaz2wsx"
ENV REDIS_DB=""
ENV SESSION_SECRET="1qaz2wsx"
ENV GIN_MODE="release"
ENV OSS_END_POINT="oss-cn-hongkong.aliyuncs.com"
ENV OSS_ACCESS_KEY_ID="LTAI4G8gfVdS9bTVBDZdWkjz"
ENV OSS_ACCESS_KEY_SECRET="kTDDj49Y9gRND2TLr1038K5LMZhUvu"
ENV OSS_BUCKET="xilixili"

#SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build main.go
COPY xili-server /usr/bin/xili-server
RUN chmod +x /usr/bin/xili-server

ENTRYPOINT ["xili-server"]

