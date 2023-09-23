# ビルド
FROM golang:1.21.0 AS builder

COPY ./ /workspace/go
WORKDIR /workspace/go

ENV ARCH="arm64"

ARG VERSION
RUN go mod download && \
    GOOS=linux GOARCH=${ARCH} CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -X main.version=${VERSION}" -o ./helloworld ./helloworld/*

# 本番実行用
# 2023/08/29 の latest
FROM gcr.io/distroless/static@sha256:2368c04cb307fd5244b92de95bd2bde6a7eb0eb4b9a0428cb276beeae127f118 as aws
COPY --from=builder ./ ./
# entrypointはlambdaの設定で上書きされる
ENTRYPOINT ["/helloworld"]

# ローカル実行用
FROM public.ecr.aws/lambda/provided:al2 as local
COPY --from=builder ./ ./
ENTRYPOINT ["/usr/local/bin/aws-lambda-rie"]
