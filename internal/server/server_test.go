package server

import (
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestServerPersist(t *testing.T) {
	file := "test.log"
	os.Remove(file)
	srv, err := NewServer(file)
	if err != nil {
		t.Fatalf("new server: %v", err)
	}
	defer srv.Close()

	req := httptest.NewRequest("POST", "/send", strings.NewReader("hi"))
	w := httptest.NewRecorder()
	srv.HandleSend(w, req)
	if w.Code != 200 {
		t.Fatalf("unexpected code: %d", w.Code)
	}

	// wait for goroutine to write to file
	time.Sleep(10 * time.Millisecond)

	data, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("read file: %v", err)
	}
	if !strings.Contains(string(data), "hi") {
		t.Fatalf("expected content in history")
	}
}
