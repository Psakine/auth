FROM golang:1.22-alpine AS builder

COPY . /github.com/Psakine/auth/source/
WORKDIR /github.com/Psakine/auth/source/

RUN go mod download
RUN go build -o ./bin/auth_server ./cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/Psakine/auth/source/bin/auth_server .

CMD ["./auth_server"]