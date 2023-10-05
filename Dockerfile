# Goの公式イメージをベースにする
FROM golang:latest

# アプリケーションのソースコードをコピー
WORKDIR /app
COPY . .

# 必要なパッケージをダウンロード
RUN go get -d -v ./...

# アプリケーションのビルド
RUN go build -o main .

EXPOSE 8080
# コンテナ内でアプリケーションを実行
CMD ["./main"]
