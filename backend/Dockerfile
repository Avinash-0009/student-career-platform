FROM golang:1.27.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server main.go


FROM alpine:3.22

RUN apk update && apk upgrade

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]