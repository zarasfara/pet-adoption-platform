package v1

import (

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func InitRoutes() *gin.Engine {
	router := gin.Default()

	return router
}
