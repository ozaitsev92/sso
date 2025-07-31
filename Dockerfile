FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=1

RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY internal ./internal
COPY cmd ./cmd

RUN go build -o /app/sso ./cmd/sso
RUN go build -o /app/migrator ./cmd/migrator

FROM alpine:3.20

RUN apk add --no-cache sqlite-libs

WORKDIR /app

COPY --from=builder /app/sso /app/sso
COPY --from=builder /app/migrator /app/migrator

EXPOSE 50051

CMD ["/app/sso"]
    