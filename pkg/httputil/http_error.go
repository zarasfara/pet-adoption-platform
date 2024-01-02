package httputil

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	EmptyResponseError = "no result for query"
)

type HTTPError struct {
	Error string `json:"error"`
}

func NewHTTPErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, HTTPError{message})
}
