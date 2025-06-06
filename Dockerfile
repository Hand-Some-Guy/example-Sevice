## 빌드 스테이션 
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 실행 스테이지
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
ENV PORT=8080
EXPOSE 8080
CMD ["./main"]