FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server main.go

EXPOSE 8081

CMD ["./server"]
