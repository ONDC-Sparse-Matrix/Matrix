FROM golang:1.21.0

WORKDIR /app1

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /backend ./main.go

EXPOSE 3000
EXPOSE 8080
# RUN ["chmod", "+x", "/app1"]
CMD ["sh", "-c", "/backend"]



