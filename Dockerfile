FROM golang:1.24-alpine AS builder

WORKDIR /api

COPY . .

RUN go mod download && \
    go build -v -o /api/api-service  ./cmd/web/main.go

FROM alpine

WORKDIR /api

COPY --from=builder /api/api-service .

EXPOSE 3000

ENTRYPOINT [ "/api/api-service"] 
