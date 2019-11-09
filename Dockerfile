FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/zhangtao25/mangostreetSerGin
COPY . $GOPATH/src/github.com/zhangtao25/mangostreetSerGin
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./mangostreetSerGin"]
