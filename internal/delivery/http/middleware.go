package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h Handler) userIdentity(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		return
	}

	token := strings.Split(authorization, " ")[1]

	userId, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set("userId", userId)

}
