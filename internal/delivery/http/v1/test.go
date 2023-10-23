package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) InitTestRoute(api *gin.RouterGroup)  {
	test := api.Group("test")
	{
		test.GET("/", h.test)
	}
}

func (h *Handler) test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}