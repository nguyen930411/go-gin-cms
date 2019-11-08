FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/nguyen930411/go-gin-cms
COPY . $GOPATH/src/github.com/nguyen930411/go-gin-cms
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]
