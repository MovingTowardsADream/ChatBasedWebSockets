package v1

import (
	"ChatBasedWebSockets/internal/entity"
	"ChatBasedWebSockets/internal/usecase"
	"context"
	"github.com/gin-gonic/gin"
)

type Authorization interface {
	CreateUser(ctx context.Context, user entity.User) (string, error)
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type authRoutes struct {
	authUseCase Authorization
}

func newAuthRoutes(handler *gin.RouterGroup, auc usecase.AuthUseCase) {
	r := &authRoutes{
		authUseCase: &auc,
	}

	handler.POST("/sign-up", r.signUp)
	handler.POST("/sign-in", r.signIn)
}

func (r *authRoutes) signUp(c *gin.Context) {

}

func (r *authRoutes) signIn(c *gin.Context) {

}
