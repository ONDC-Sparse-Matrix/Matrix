FROM golang:1.21.0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /centralcdn ./pkg/cmd/main.go

EXPOSE 3001
EXPOSE 8080

CMD ["sh", "-c", "/centralcdn"]

