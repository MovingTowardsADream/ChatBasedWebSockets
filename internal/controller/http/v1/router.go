package v1

import (
	"ChatBasedWebSockets/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
	"net/http"
)

func NewRouter(handler *gin.Engine, l *slog.Logger, auc *usecase.AuthUseCase) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	auth := handler.Group("/auth")
	{
		newAuthRoutes(auth, auc)
	}

	// Authorization
	authMiddleware := &AuthMiddleware{
		auc,
	}

	h := handler.Group("/api/v1", authMiddleware.UserIdentity())
	{
		newUsersRoutes(h)
	}
}
