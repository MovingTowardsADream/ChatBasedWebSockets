package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	userIdCtx = "userId"
)

type AuthMiddleware struct {
	authUseCase Authorization
}

func (h *AuthMiddleware) UserIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := bearerToken(c.Request)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userId, err := h.authUseCase.ParseToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(userIdCtx, userId)

		c.Next()
	}
}

func bearerToken(r *http.Request) (string, bool) {
	const prefix = "Bearer "

	header := r.Header.Get("Authorization")
	if header == "" {
		return "", false
	}

	if len(header) > len(prefix) && strings.EqualFold(header[:len(prefix)], prefix) {
		return header[len(prefix):], true
	}

	return "", false
}

func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get(userIdCtx)
	if !ok {
		return "", errors.New("user id not found")
	}

	idString, ok := id.(string)
	if !ok {
		return "", errors.New("user id is of invalid type")
	}

	return idString, nil
}
