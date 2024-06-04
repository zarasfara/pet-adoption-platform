package http

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/zarasfara/pet-adoption-platform/docs"
	v1 "github.com/zarasfara/pet-adoption-platform/internal/delivery/http/v1"
	"github.com/zarasfara/pet-adoption-platform/internal/service"
)

type handler struct {
	services *service.Services
	appURL   string
	refreshTokenTTL time.Duration
}

func NewHandler(
	services *service.Services, 
	appURL string, 
	refreshTokenTTL time.Duration,
	) *handler {
	return &handler{
		services: services,
		appURL: appURL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

// Init инициализирует роутер и прикрепляет маршруты
func (h handler) Init() *gin.Engine {
	router := gin.Default()

	h.initAPI(router)

	return router
}

// initAPI инициализирует route группу api
func (h handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)

	auth := router.Group("/auth")
	{
		h.InitAuthRoutes(auth)
	}

	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
