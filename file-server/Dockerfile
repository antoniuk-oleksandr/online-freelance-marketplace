FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/file-server ./cmd/file-server

CMD ["/app/bin/file-server"]
