package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
)

func (h Handler) InitAuthRoutes(auth *gin.RouterGroup) {
	auth.POST("/sign-up", h.signUp)
	auth.POST("/sign-in", h.signIn)
}

func (h *Handler) signUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.services.Authorization.CreateUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func (h Handler) signIn(c *gin.Context) {

	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := h.services.Authorization.GenerateToken(body.Email, body.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}
