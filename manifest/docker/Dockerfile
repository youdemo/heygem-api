# 第一阶段：编译 Go 服务
FROM golang:1.24.1-alpine3.21 AS builder

WORKDIR /app

# 复制 go.mod 和 go.sum 文件并下载依赖
COPY . .
RUN go mod tidy

# 复制源码并编译
RUN go build -o ./linux_amd64/heygem-api main.go


# 从 builder 阶段复制编译好的二进制文件
FROM alpine:latest

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR                   /app
ENV APP_NAME                  heygem-api
ADD manifest/config/config.docker.yaml $WORKDIR/config.yaml
COPY --from=builder /app/linux_amd64/$APP_NAME $WORKDIR/$APP_NAME
RUN chmod +x $WORKDIR/$APP_NAME
RUN apk update && apk add --no-cache yasm ffmpeg && rm -rf /var/cache/apk/*


###############################################################################
#                                   START
###############################################################################
EXPOSE 8000
WORKDIR $WORKDIR
CMD ./$APP_NAME
