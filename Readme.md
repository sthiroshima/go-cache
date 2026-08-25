# go-cache

A simple Redis-like in-memory cache written in Go.

This is a pet project for learning Go and understanding how a basic Redis server works under the hood.

## Features

* TCP server
* RESP protocol
* Multiple client connections
* Concurrent access to storage
* Graceful shutdown
* In-memory storage

## Supported commands

```text
SET key value
GET key
DEL key
EXISTS key
KEYS *
DBSIZE
FLUSHDB
```

## Run

```bash
go run ./cmd
```

The server listens on port `6379`.

You can connect with `redis-cli`:

```bash
redis-cli -p 6379
```

Example:

```text
SET name go-cache
GET name
EXISTS name
DEL name
```

## Project structure

```text
cmd/
    main.go

internal/
    handler.go
    parser.go
    protocol.go
    server.go
    storage.go
```

## What I wanted to practice

* TCP connections in Go
* goroutines
* mutexes and concurrent access to maps
* RESP parsing
* graceful shutdown
* working with `context` and OS signals

## Docker

Build and start the server:
```text
docker compose up -d --build
```

Stop the server:
```text
docker compose down
```
View logs:
```text
docker compose logs -f
```
The server is available on port 6379.

## Version

`v1.1.0`
