package server

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"sync"
)

// Server represents the chat server using Server-Sent Events.
type Server struct {
	mu         sync.RWMutex
	clients    map[int]chan string
	broadcast  chan string
	register   chan chan string
	unregister chan chan string
	history    *os.File
	nextID     int
}

// NewServer creates a new chat server.
func NewServer(historyFile string) (*Server, error) {
	f, err := os.OpenFile(historyFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	s := &Server{
		clients:    make(map[int]chan string),
		broadcast:  make(chan string, 256),
		register:   make(chan chan string),
		unregister: make(chan chan string),
		history:    f,
	}

	go s.run()

	return s, nil
}

func (s *Server) run() {
	for {
		select {
		case ch := <-s.register:
			s.mu.Lock()
			s.nextID++
			s.clients[s.nextID] = ch
			s.mu.Unlock()
		case ch := <-s.unregister:
			s.mu.Lock()
			for id, c := range s.clients {
				if c == ch {
					delete(s.clients, id)
					close(c)
					break
				}
			}
			s.mu.Unlock()
		case msg := <-s.broadcast:
			s.mu.RLock()
			for _, ch := range s.clients {
				select {
				case ch <- msg:
				default:
				}
			}
			s.mu.RUnlock()
			if _, err := s.history.WriteString(msg + "\n"); err != nil {
				log.Printf("failed to persist message: %v", err)
			}
		}
	}
}

// HandleEvents handles SSE connections.
func (s *Server) HandleEvents(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")

	ch := make(chan string, 10)
	s.register <- ch
	defer func() { s.unregister <- ch }()

	for msg := range ch {
		_, _ = w.Write([]byte("data: " + msg + "\n\n"))
		flusher.Flush()
	}
}

// HandleSend handles POST messages.
func (s *Server) HandleSend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := bufio.NewReader(r.Body).ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	s.broadcast <- body
}

// Close closes the server's history file.
func (s *Server) Close() error {
	return s.history.Close()
}
