# Live Chat Server

This project provides a simple Node.js based chat server that can be used as a
starting point for integrating live streaming chat features. It is designed to
be extended with third‑party messaging services such as **RongCloud** and
**Tencent IM**.

## Features

- WebSocket based real-time messaging using `socket.io`.
- Basic HTTP endpoint to verify the server is running.
- Configuration file placeholder (`config.json`) for third-party credentials.

## Getting Started

1. Copy `config.example.json` to `config.json` and fill in your credentials for
   RongCloud or Tencent IM.
2. Install dependencies (requires internet access):

```bash
npm install express socket.io
```

3. Start the server:

```bash
node server.js
```

The server listens on port `3000` by default and exposes a WebSocket endpoint
for clients to exchange messages.

## Notes

Actual API calls to RongCloud or Tencent IM are not implemented in this sample.
You can add them inside the `chat message` handler in `server.js`.
