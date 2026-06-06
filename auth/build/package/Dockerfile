# Build
FROM golang:1.26.1-alpine3.23 AS builder
WORKDIR /app
ADD go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth.exe ./cmd/service/*.go

#Run
FROM alpine:3.23
WORKDIR /app
# COPY --from=builder /app/config/.env .
COPY --from=builder /app/auth.exe .
CMD ["./auth.exe"]