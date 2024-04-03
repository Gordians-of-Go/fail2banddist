package main

import (
	"encoding/json"
	"flag"
	"log/slog"
	"net/url"

	"github.com/Gordians-of-Go/fail2banddist/pkg/response"
	"github.com/gorilla/websocket"
)

var port string
var upgrader = websocket.Upgrader{}

func init() {
	flag.StringVar(&port, "p", "8080", "Port to listen on")
	flag.Parse()
}

type client struct {
}

func main() {
	serverWsURL := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/hello",
	}
	c, _, err := websocket.DefaultDialer.Dial(serverWsURL.String(), nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	var d response.Dummy
	for {

		_, message, err := c.ReadMessage()
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(message, &d)
		if err != nil {
			slog.Info("Failed to unmarshal message", "error", err)
			continue
		}
		slog.Info("Received message", "message", d.Message)
	}
}
