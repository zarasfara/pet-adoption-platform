package http

import (
	"github.com/zarasfara/pet-adoption-platform/pkg/httputil"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
)

func (h Handler) InitAuthRoutes(auth *gin.RouterGroup) {
	auth.POST("/sign-up", h.signUp)
	auth.POST("/sign-in", h.signIn)
	auth.GET("/current-user", h.userIdentity, h.getCurrentUser)
}

// @Summary	Регистрация
// @Tags		auth
// @Accept		json
// @Param		model	body	models.User	true	"Регистрация пользователя"
// @Accept		json
// @Success	204
// @Failure	400	{object}	httputil.HTTPError
// @Failure	500	{object}	httputil.HTTPError
// @Router		/auth/sign-up [post]
func (h Handler) signUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Authorization.CreateUser(user)
	if err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

type tokenResponse struct {
	AccessToken string `json:"accessToken"`
}

// @Summary	Аутентификация
// @Tags		auth
// @Accept		json
// @Produce	json
// @Param		model	body		http.signIn.signInInput				true	"Аутентификация пользователя"
// @Success	200		{object}	tokenResponse{accessToken=string}	"accessToken"
// @Failure	400		{object}	httputil.HTTPError
// @Failure	500		{object}	httputil.HTTPError
// @Router		/auth/sign-in [post]
func (h Handler) signIn(c *gin.Context) {
	type signInInput struct {
		Email    string `json:"email" binding:"required" format:"email"`
		Password string `json:"password" binding:"required" format:"password"`
	}
	var body signInInput

	if err := c.BindJSON(&body); err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(body.Email, body.Password)
	if err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponse{AccessToken: token})
}

// @Summary	Получить текущего пользователя
// @Tags		auth
// @Produce	json
// @Security	BearerAuth
// @Success	200	{object}	models.User	"Текущий пользователь"
// @Failure	400	{object}	httputil.HTTPError
// @Failure	500	{object}	httputil.HTTPError
// @Router		/auth/current-user [get]
func (h Handler) getCurrentUser(c *gin.Context) {
	userId := c.GetInt("userId")
	if userId == 0 {
		httputil.NewHTTPErrorResponse(c, http.StatusBadRequest, "not authenticated")
		return
	}

	user, err := h.services.Authorization.GetCurrentUser(userId)
	if err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
