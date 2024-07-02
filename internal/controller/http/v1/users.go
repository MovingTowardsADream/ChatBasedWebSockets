package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type userRoutes struct{}

func newUsersRoutes(handler *gin.RouterGroup) {
	r := &userRoutes{}

	handler.POST("/test", r.TestAuth)
}

func (r *userRoutes) TestAuth(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	fmt.Println(userId)
}
