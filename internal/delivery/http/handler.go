package http

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zarasfara/pet-adoption-platform/internal/delivery/http/v1"
)

type Handler struct {
}

func NewHandler() *Handler  {
	return new(Handler)
}

// Init инициализирует роутер и прикрепляет маршруты
func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	h.initAPI(router)

	return router
}

// initAPI инициализирует route группу api
func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler()

	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}