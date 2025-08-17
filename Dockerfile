FROM golang:1.23-alpine as builder
WORKDIR /app
COPY app/go.mod app/go.sum ./
RUN go mod download
COPY /app .
RUN go build -o main .
RUN chmod +x ./main

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/public ./public
EXPOSE 8080
CMD ["./main"]