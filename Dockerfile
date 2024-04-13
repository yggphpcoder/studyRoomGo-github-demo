FROM golang:1.20.0-alpine3.17 as builder


ENV GO111MODULE=on     GOPROXY=https://goproxy.cn,direct



RUN mkdir /app

WORKDIR /app

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download


ADD . /app/



RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o main ./cmd/studyRoomGo

FROM alpine:latest

# 容器默认时区为UTC，如需使用上海时间请启用以下时区设置命令
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo Asia/Shanghai > /etc/timezone
# grpc
RUN apk add grpc
#protobuf
RUN apk add  protobuf-dev

WORKDIR /app

COPY --from=builder /app/main .

ADD ./configs /data/conf


EXPOSE 8080
EXPOSE 61005/udp
VOLUME /data/conf

CMD ["/app/main", "-conf", "/data/conf/config.pro.yaml"]


#docker build --no-cache=true -t study-room-go  .
#docker build -t study-room-go  .
#docker tag study-room-go ccr.ccs.tencentyun.com/tcb-100027762657-rvcd/study-room-go:latest
#docker push ccr.ccs.tencentyun.com/tcb-100027762657-rvcd/study-room-go:latest
#docker pull ccr.ccs.tencentyun.com/tcb-100027762657-rvcd/study-room-go:latest
