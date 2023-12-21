package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zarasfara/pet-adoption-platform/pkg/httputil"
)

func (h Handler) userIdentity(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		httputil.NewHTTPErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	token, err := extractTokenFromAuthorizationHeader(authorization)
	if err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	userId, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Use userId as needed
	c.Set("userId", userId)
}

func extractTokenFromAuthorizationHeader(authorization string) (string, error) {
	parts := strings.Split(authorization, " ")
	if len(parts) != 2 {
		return "", errors.New("invalid auth header format")
	}
	return parts[1], nil
}
