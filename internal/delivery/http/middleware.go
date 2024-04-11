package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zarasfara/pet-adoption-platform/pkg/httputil"
)

var ErrInvalidAuthHeaderFormat = errors.New("invalid auth header format")

func (h handler) userIdentity(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		httputil.NewHTTPErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	accessToken, err := extractToken(authorizationHeader)
	if err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.services.Authorization.UserIDFromToken(accessToken)
	if err != nil {
		// token is expires or something else
		httputil.NewHTTPErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", userId)
}

func extractToken(authorizationHeader string) (string, error) {
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 {
		return "", ErrInvalidAuthHeaderFormat
	}
	return parts[1], nil
}
