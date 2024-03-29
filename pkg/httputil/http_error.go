package httputil

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HTTPError struct {
	ErrorMessage string `json:"error"`
}

func NewHTTPErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, HTTPError{message})
}
