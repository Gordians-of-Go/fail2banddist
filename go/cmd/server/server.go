package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Gordians-of-Go/fail2banddist/pkg/server"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var port string
var upgrader = websocket.Upgrader{}

func init() {
	flag.StringVar(&port, "p", "8080", "Port to listen on")
	flag.Parse()
}

func main() {
	// Handle signals
	server, err := server.NewServer(port)
	if err != nil {
		slog.Info("Failed to create server", "error", err)
		os.Exit(1)
	}
	stopper := make(chan os.Signal, 1)
	signal.Notify(stopper, os.Interrupt, syscall.SIGTERM)

	r := mux.NewRouter()
	r.HandleFunc("/hello", server.HandleNewClient)
	http.Handle("/", r)
	listenAddress := fmt.Sprintf(":%s", port)
	go http.ListenAndServe(listenAddress, nil)
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		c := 0
		for {
			select {
			case <-ticker.C:
				server.NotifyClients(fmt.Sprintf("Hello %d", c))
				c++
			}
		}
	}()

	<-stopper
	slog.Info("Server stopped")
}
