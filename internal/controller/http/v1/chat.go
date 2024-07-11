package v1

import (
	"ChatBasedWebSockets/internal/controller/http/v1/ws"
	"github.com/gin-gonic/gin"
)

type chatRoutes struct {
	manager *ws.Manager
}

func newChatRoutes(handler *gin.RouterGroup, m *ws.Manager) {
	r := &chatRoutes{manager: m}

	handler.GET("/chat", r.Chat)
}

func (r *chatRoutes) Chat(c *gin.Context) {
	r.manager.ServeWs(c.Writer, c.Request)
}
