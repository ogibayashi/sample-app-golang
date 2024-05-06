FROM golang:1.21.6 as builder
WORKDIR /app
COPY . .

RUN make build

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/bin/sample-app-golang .
COPY --from=builder /app/deploy/ .

CMD ["./sample-app-golang"]
