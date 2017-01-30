package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(serveWs))
	defer ts.Close()

	url := makeWsProto(ts.URL)

	// u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}

	log.Printf("connecting to %s", url)

	c, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		log.Fatal("dial:", err)
	}

	defer c.Close()

	err = c.WriteMessage(websocket.TextMessage, []byte("hi!"))

	if err != nil {
		t.Error("write:", err)
	}

	_, msg, err := c.ReadMessage()

	if err != nil {
		t.Error("read:", err, msg)
	}

	if string(msg) != "hi!" {
		t.Error("expected 'hi!', got", err, string(msg))

	}
}

func makeWsProto(s string) string {
	return "ws" + strings.TrimPrefix(s, "http")
}
