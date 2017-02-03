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

	c1, _, err1 := websocket.DefaultDialer.Dial(url, nil)
	c2, _, err2 := websocket.DefaultDialer.Dial(url, nil)

	if err1 != nil {
		log.Fatal("dial:", err1)
	}

	if err2 != nil {
		log.Fatal("dial:", err2)
	}

	defer c1.Close()
	defer c2.Close()

	testMessage := []byte("hi from client1!")

	err1 = c1.WriteMessage(websocket.TextMessage, testMessage)

	if err1 != nil {
		t.Error("write:", err1)
	}

	_, msg1, err1 := c1.ReadMessage()

	if err1 != nil {
		t.Error("read:", err1, msg1)
	}

	if string(msg1) != string(testMessage) {
		t.Error("expected %s, got %s", err1, string(msg1))
	}

	_, msg2, err2 := c2.ReadMessage()

	if err2 != nil {
		t.Error("read:", err2, msg2)
	}

	if string(msg2) != string(testMessage) {
		t.Error("expected %s, got %s", err1, string(msg2))
	}
}

func makeWsProto(s string) string {
	return "ws" + strings.TrimPrefix(s, "http")
}
