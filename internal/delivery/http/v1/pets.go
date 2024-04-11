package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zarasfara/pet-adoption-platform/pkg/httputil"
	"net/http"
)

func (h Handler) initPetsRoutes(rg *gin.RouterGroup) {
	pets := rg.Group("/pets")
	{
		pets.GET("", h.getAllPets)
	}
}

//	@Summary		Get all pets
//	@Tags			pets
//	@Description	Retrieves all pets
//	@Produce		json
//	@Param			sortField	query		string	false	"Sort field"
//	@Success		200			{array}		models.Pet
//	@Failure		404			{array}		models.Pet
//	@Failure		500			{object}	httputil.HTTPError
//	@Router			/api/v1/pets [get]
func (h Handler) getAllPets(c *gin.Context) {
	sortField := c.Query("sortField")

	pets, err := h.services.Pet.PetsBySortField(sortField)
	if err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(pets) == 0 {
		c.JSON(http.StatusNotFound, pets)
		return
	}

	c.JSON(http.StatusOK, pets)
}
