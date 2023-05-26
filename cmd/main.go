package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zarasfara/pet-adoption-platoform/internal/routes"
)

func main() {

	router := gin.Default()

	router.GET("test", routes.Test)

	router.Run("127.0.0.1:8000")

}
