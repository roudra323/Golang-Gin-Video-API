package api

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/roudra323/gin-simple-api/service"
)

type SocketHandler struct {
	service service.VideoService
}

func NewSocketHandler(s service.VideoService) *SocketHandler {
	return &SocketHandler{service: s}
}

func (h *SocketHandler) RegisterHandlers(server *socketio.Server) {

	// The OnConnect handler is fired when a new client connects
	server.OnConnect("/", func(c socketio.Conn) error {
		c.SetContext("")
		log.Printf("New client connected: %s", c.ID())
		return nil
	})

	// The OnEvent handler is fired for a specific custom event.
	// Here, we listen for a "get_videos" event from any client.
	server.OnEvent("/", "get_videos", func(c socketio.Conn) {
		log.Printf("Received 'get_videos' event from client %s", c.ID())

		// We call our existing service to get the data
		videos, err := h.service.GetAll()

		if err != nil {

			c.Emit("get_videos_error", err.Error())
			return
		}

		c.Emit("get_videos_response", videos)
	})

	// The OnError handler is fired for any connection error.
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Printf("Socket Error for client %s: %v", s.ID(), e)
	})

	// The OnDisconnect handler is fired when a client disconnects.
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Printf("Client %s disconnected for reason: %s", s.ID(), reason)
	})
}
