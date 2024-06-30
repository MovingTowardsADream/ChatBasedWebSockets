package v1

import (
	"github.com/gin-gonic/gin"
)

type userRoutes struct{}

func newUsersRoutes(handler *gin.RouterGroup) {
	r := &userRoutes{}

	handler.POST("/test", r.TestId)
}

func (r *userRoutes) TestId(c *gin.Context) {

}
