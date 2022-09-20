# バイナリ作成（本番環境用ビルド実行）コンテナ
FROM golang:1.19.1-bullseye as release-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# -----------------------------------------------------------

# 本番実行環境コンテナ
FROM debian:bullseye-slim as release

RUN apt-get update

COPY --from=release-builder /app/app .

CMD ["./app"]

# -----------------------------------------------------------

# ローカル実行環境コンテナ（ホットリロード環境）
FROM golang:1.19.1 as dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
