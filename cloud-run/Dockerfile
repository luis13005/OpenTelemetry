FROM golang:alpine as builder

RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS="linux" CGO_ENABLED=0 go build -o server cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/server .
COPY --from=builder /app/cmd/.env cmd/.env

EXPOSE 8080

ENTRYPOINT ["/app/server"]