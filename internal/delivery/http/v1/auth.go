package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
)

func (h *Handler) InitAuthRoutes(api *gin.RouterGroup) {
	test := api.Group("auth")
	{
		test.POST("/sign-up", h.signUp)
	}
}

func (h *Handler) signUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Authorization.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
