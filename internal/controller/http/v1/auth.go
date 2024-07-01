package v1

import (
	"ChatBasedWebSockets/internal/entity"
	"ChatBasedWebSockets/internal/usecase"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Authorization interface {
	CreateUser(ctx context.Context, user entity.User) (string, error)
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type authRoutes struct {
	authUseCase Authorization
}

func newAuthRoutes(handler *gin.RouterGroup, auc *usecase.AuthUseCase) {
	r := &authRoutes{
		authUseCase: auc,
	}

	handler.POST("/sign-up", r.signUp)
	handler.POST("/sign-in", r.signIn)
}

type signUpInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *authRoutes) signUp(c *gin.Context) {
	var input signUpInput

	if err := c.Bind(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// TODO Validate input data

	user := entity.User{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	id, err := r.authUseCase.CreateUser(c.Request.Context(), user)

	if err != nil {
		// TODO Handler error

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	type response struct {
		Id string `json:"id"`
	}

	c.JSON(http.StatusCreated, response{
		Id: id,
	})
}

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *authRoutes) signIn(c *gin.Context) {
	var input signInInput

	if err := c.Bind(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// TODO Validate

	token, err := r.authUseCase.GenerateToken(c.Request.Context(), input.Username, input.Password)

	if err != nil {
		// TODO Handler error

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	type response struct {
		Token string `json:"token"`
	}

	c.JSON(http.StatusOK, response{
		Token: token,
	})
}
