FROM golang:1.21.6 as builder
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN make build

FROM alpine:latest
WORKDIR /root/
RUN mkdir deploy
COPY --from=builder /app/bin/sample-app-golang .
COPY --from=builder /app/deploy/ ./deploy/

CMD ["./sample-app-golang"]
