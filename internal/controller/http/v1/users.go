package v1

import (
	"ChatBasedWebSockets/internal/entity"
	"ChatBasedWebSockets/internal/usecase"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type UsersWork interface {
	GetAllUsers(ctx context.Context, userCh chan<- entity.User)
}

type userRoutes struct {
	usersUseCase UsersWork
}

func newUsersRoutes(handler *gin.RouterGroup, uuc *usecase.UsersUseCase) {
	r := &userRoutes{usersUseCase: uuc}

	handler.GET("/users", r.GetAllUsers)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (r *userRoutes) GetAllUsers(c *gin.Context) {
	usersCh := make(chan entity.User)

	r.usersUseCase.GetAllUsers(c.Request.Context(), usersCh)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	defer conn.Close()

	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}

	for user := range usersCh {
		conn.WriteJSON(user)
	}
}
