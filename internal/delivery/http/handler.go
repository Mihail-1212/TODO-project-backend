package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mihail-1212/todo-project-backend/docs"
	v1 "github.com/mihail-1212/todo-project-backend/internal/delivery/http/v1"
	"github.com/mihail-1212/todo-project-backend/internal/service"
	"github.com/mihail-1212/todo-project-backend/pkg/auth"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services   *service.Services
	authorizer *auth.Authorizer
}

func NewHandler(services *service.Services, authorizer *auth.Authorizer) *Handler {
	return &Handler{
		services:   services,
		authorizer: authorizer,
	}
}

func (h *Handler) InitAPI() *gin.Engine {
	router := gin.New()
	handlerV1 := v1.NewHandler(h.services, h.authorizer)

	router.Use(optionMiddleware)
	// Added cors headers
	router.Use(corsMiddleware)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}

	return router
}
