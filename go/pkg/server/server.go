package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Gordians-of-Go/fail2banddist/pkg/response"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// Server holds the server configuration
type Server struct {
	Port        string
	connections []*websocket.Conn
}

// HandleNewClient handles new client connections
func (s *Server) HandleNewClient(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Info("Failed to upgrade connection", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s.connections = append(s.connections, ws)
	slog.Info("New client connected", "client", ws.RemoteAddr())
}

func (s *Server) NotifyClients(msg string) {
	message := response.Dummy{
		Message: msg,
	}
	encodedMessage, err := json.Marshal(&message)
	if err != nil {
		panic(err)
	}
	for _, c := range s.connections {

		err := c.WriteMessage(websocket.TextMessage, encodedMessage)
		if err != nil {
			slog.Info("Failed to write message", "error", err, "client", c.RemoteAddr())
		}
	}
}

// NewServer creates a server with the given port
func NewServer(port string) (*Server, error) {
	p, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	s := &Server{
		Port: fmt.Sprintf("%d", p),
	}
	return s, nil
}
