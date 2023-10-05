FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main /app/server/main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 50051
CMD [ "/app/main" ]
