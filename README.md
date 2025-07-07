# im

A lightweight Go-based instant messaging server. The implementation uses Server-Sent Events (SSE) and an in-memory queue to broadcast messages to all connected clients. Messages are persisted to a local log file. The server can handle at least 100 concurrent clients using Go's goroutines and channels.

## Running locally

```
go run ./cmd/server
```

The server listens on `:8080` with the following endpoints:

- `POST /send` - send a message
- `GET  /events` - receive a stream of messages via SSE

## Features

- Broadcast messages to all connected clients.
- In-memory caching of active connections.
- Simple queue using Go channels for message distribution.
- Persistence of message history to a file.
- Ready for load balancing by running multiple instances behind a reverse proxy.

## GitHub Actions

A basic CI workflow builds and tests the project on each push.

