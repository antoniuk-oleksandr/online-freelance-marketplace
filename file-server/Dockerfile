FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod ./
run go mod download

COPY . .

RUN go build -o /go-file-server main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /go-file-server ./
COPY uploads/ ./uploads/

EXPOSE 8030

CMD ["./go-file-server"]
