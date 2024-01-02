package v1

import (
	"github.com/gin-gonic/gin"
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

// Init инициализирует группу v1 с маршрутами приложения
func (h Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initPetsRoutes(v1)
	}
}
