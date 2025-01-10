FROM golang:1.21-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /app/main .

FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

RUN adduser -D -g '' appuser

WORKDIR /app

RUN mkdir -p /app/logs && \
    chown -R appuser:appuser /app

COPY --from=builder /app/main .
COPY .env .

COPY --chown=appuser:appuser config.yml .
COPY --chown=appuser:appuser translation/ ./translation/

USER appuser

EXPOSE 15340

CMD ["./main"]