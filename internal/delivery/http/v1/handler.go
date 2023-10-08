package v1

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return new(Handler)
}

// Init инициализует группу v1 с машрутами приложения
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.InitTestRoute(v1)
	}
}
