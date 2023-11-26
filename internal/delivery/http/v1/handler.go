package v1

import (
	"net/http"

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

// Init инициализует группу v1 с машрутами приложения
func (h Handler) Init(api *gin.RouterGroup) {	
	v1 := api.Group("/v1")
	{
		v1.GET("/ping", h.ping)
	}
}

func (h Handler) ping(c *gin.Context) {
	id, _ := c.Get("userId") 

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"user": id,
	})
}
