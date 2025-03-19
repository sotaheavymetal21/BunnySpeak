FROM golang:1.20-alpine AS builder

WORKDIR /app

# Go Moduleの依存関係をコピーしてインストール
COPY go.mod go.sum* ./
RUN go mod download

# ソースコードをコピー
COPY . .

# バイナリをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bunnyspeak .

# 最終的なイメージを作成
FROM alpine:latest  

WORKDIR /root/

# ビルドされたバイナリをコピー
COPY --from=builder /app/bunnyspeak .

# エントリーポイントを設定
ENTRYPOINT ["./bunnyspeak"] 