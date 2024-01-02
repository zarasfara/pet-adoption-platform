package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zarasfara/pet-adoption-platform/pkg/httputil"
	"net/http"
)

func (h Handler) initPetsRoutes(rg *gin.RouterGroup) {
	pets := rg.Group("/pets")
	{
		pets.GET("/", h.getAllPets)
	}
}

func (h Handler) getAllPets(c *gin.Context) {
	sortField := c.Query("sortField")

	pets, err := h.services.Pet.GetAll(sortField)
	if err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if len(pets) == 0 {
		httputil.NewHTTPErrorResponse(c, http.StatusNotFound, httputil.EmptyResponseError)
	}

	c.JSON(http.StatusOK, pets)
}
