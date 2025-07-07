package main

import (
	"log"
	"net/http"

	"im/internal/server"
)

func main() {
	srv, err := server.NewServer("history.log")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	defer srv.Close()

	http.HandleFunc("/events", srv.HandleEvents)
	http.HandleFunc("/send", srv.HandleSend)

	log.Println("server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("listen: %v", err)
	}
}
