# 使用适合Go应用的基础镜像
FROM golang:1.22-alpine as builder
RUN apk update && apk add --no-cache upx && rm -rf /var/cache/apk/*
RUN apk add make

# 设置工作目录
WORKDIR /app

# 复制所有文件到容器中
COPY . .

# 下载依赖
RUN go mod tidy

# 构建应用程序
#RUN CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -o deeplx .
RUN make build-all
RUN upx -6 build/oaiproxy_linux_*

FROM scratch as final
ARG TARGETOS
ARG TARGETARCH
WORKDIR /home/oai-proxy/
COPY --from=builder /app/build/oaiproxy_${TARGETOS}_${TARGETARCH} /home/oai-proxy/oaiproxy
COPY --from=builder /app/config.yml /home/oai-proxy/config.yml
COPY --from=builder /app/template/* /home/oai-proxy/template/

# 开放端口
EXPOSE 8080

# 运行
CMD ["./oaiproxy"]