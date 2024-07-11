package v1

import (
	"ChatBasedWebSockets/internal/controller/http/v1/ws"
	"ChatBasedWebSockets/internal/entity"
	"ChatBasedWebSockets/internal/usecase"
	"context"
	"github.com/gin-gonic/gin"
	"log"
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
	handler.GET("/chatWs", r.Chat)
}

func (r *userRoutes) GetAllUsers(c *gin.Context) {
	usersCh := make(chan entity.User)

	r.usersUseCase.GetAllUsers(c.Request.Context(), usersCh)

	conn, err := ws.Upgrader.Upgrade(c.Writer, c.Request, nil)
	defer conn.Close()

	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}

	for user := range usersCh {
		conn.WriteJSON(user)
	}
}

func (r *userRoutes) Chat(c *gin.Context) {
	//userId, err := getUserId(c)
	//if err != nil {
	//	return
	//}
	//fmt.Println(userId)

}
