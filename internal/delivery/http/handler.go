package http

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/zarasfara/pet-adoption-platform/docs"
	v1 "github.com/zarasfara/pet-adoption-platform/internal/delivery/http/v1"
	"github.com/zarasfara/pet-adoption-platform/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

// Init инициализирует роутер и прикрепляет маршруты
func (h Handler) Init() *gin.Engine {
	router := gin.Default()

	h.initAPI(router)

	return router
}

// initAPI инициализирует route группу api
func (h Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)

	auth := router.Group("/auth")
	{
		h.InitAuthRoutes(auth)
	}

	api := router.Group("/api", h.userIdentity)
	{
		handlerV1.Init(api)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
