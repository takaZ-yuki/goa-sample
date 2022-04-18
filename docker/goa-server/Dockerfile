FROM golang:1.17-alpine3.13 AS builder
WORKDIR /app
COPY go.mod go.sum /
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/goa_sample

FROM alpine:3.10.1
WORKDIR /app
EXPOSE 8088
COPY --from=builder /app/main .
CMD ["./main"]