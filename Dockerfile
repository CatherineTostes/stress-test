FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o stress-test

FROM alpine:latest
COPY --from=builder /app/stress-test .
ENTRYPOINT ["./stress-test"]