FROM golang:1.26.1 AS builder

WORKDIR /cache

# add go.sum when will be
COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o cache ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates busybox-extras

RUN adduser -D appuser
USER appuser

WORKDIR /home/appuser

COPY --from=builder /cache .

ENV PORT=6379

EXPOSE $PORT

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD nc -z localhost $PORT || exit 1

ENTRYPOINT ["./cache"]