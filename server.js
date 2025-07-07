const express = require('express');
const http = require('http');
const { Server } = require('socket.io');
const fs = require('fs');

// Load configuration for third-party services if available
let config = {};
try {
  config = JSON.parse(fs.readFileSync('./config.json', 'utf-8'));
} catch (err) {
  console.warn('config.json not found, using default config');
}

const app = express();
const server = http.createServer(app);
const io = new Server(server, {
  cors: {
    origin: '*',
  },
});

app.get('/', (req, res) => {
  res.send('Live chat server running');
});

io.on('connection', (socket) => {
  console.log('a user connected');
  socket.on('chat message', (msg) => {
    // Forward message to 3rd-party services as needed
    // e.g., RongCloud or Tencent IM API calls (not implemented here)
    io.emit('chat message', msg);
  });
  socket.on('disconnect', () => {
    console.log('user disconnected');
  });
});

const PORT = process.env.PORT || 3000;
server.listen(PORT, () => {
  console.log(`Server listening on port ${PORT}`);
});
