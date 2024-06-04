package http

import (
	"net/http"

	"github.com/zarasfara/pet-adoption-platform/pkg/httputil"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
)

const refreshTokenKey = "refreshToken"

func (h handler) InitAuthRoutes(auth *gin.RouterGroup) {
	auth.POST("/sign-up", h.signUp)
	auth.POST("/sign-in", h.signIn)
	auth.POST("/refresh-tokens", h.refreshTokens)
	auth.GET("/current-user", h.userIdentity, h.getCurrentUser)
}

// @Summary	Регистрация
// @Tags		auth
// @Accept		json
// @Param		model	body	models.AddRecordUser	true	"Регистрация пользователя"
// @Accept		json
// @Success	204
// @Failure	400	{object}	httputil.HTTPError
// @Failure	500	{object}	httputil.HTTPError
// @Router		/auth/sign-up [post]
func (h handler) signUp(c *gin.Context) {
	var user models.AddRecordUser

	if err := c.BindJSON(&user); err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Authorization.CreateUser(user)
	if err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

type tokenResponse struct {
	AccessToken string `json:"accessToken"`
}

// @Summary	Аутентификация
// @Tags		auth
// @Accept		json
// @Produce	json
// @Param		model	body		http.signIn.signInInput	true	"Аутентификация пользователя"
// @Success	200		{object}	tokenResponse
// @Failure	400		{object}	httputil.HTTPError
// @Failure	500		{object}	httputil.HTTPError
// @Router		/auth/sign-in [post]
func (h handler) signIn(c *gin.Context) {
	type signInInput struct {
		Email    string `json:"email" binding:"required" format:"email"`
		Password string `json:"password" binding:"required" format:"password"`
	}
	var body signInInput

	if err := c.BindJSON(&body); err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.services.Authorization.SignIn(body.Email, body.Password)
	if err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie(refreshTokenKey, refreshToken,
		int(h.refreshTokenTTL.Seconds())+60, // Переделать 
		"/auth",
		h.appURL,
		false,
		true,
	)

	c.JSON(http.StatusOK, tokenResponse{AccessToken: accessToken})
}

// @Summary	Обновление пары токенов
// @Description	Обновляет пару токенов (access и refresh) на основе предоставленного access токена.
// @Tags		auth
// @Accept		json
// @Produce		json
// @Security		BearerAuth
// @Success		200		{object}	tokenResponse	"Возвращает новую пару токенов"
// @Failure		400		{object}	httputil.HTTPError	"Не удалось обновить токены"
// @Router		/auth/refresh-tokens [post]
func (h handler) refreshTokens(c *gin.Context) {
	refreshToken, err := c.Cookie(refreshTokenKey)
	if err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newAccessToken, newRefreshToken, err := h.services.Authorization.RegenerateTokens(refreshToken)
	if err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// delete old refresh token
	c.SetCookie(refreshTokenKey, "", -1, "/auth", h.appURL, false, true)

	c.SetCookie(refreshTokenKey, newRefreshToken, int(h.refreshTokenTTL.Seconds())+60, "/auth", h.appURL, false, true)

	c.JSON(http.StatusOK, gin.H{
		"accessToken": newAccessToken,
	})
}

// @Summary	Получить текущего пользователя
// @Tags		auth
// @Produce	json
// @Security	BearerAuth
// @Success	200	{object}	models.User	"Текущий пользователь"
// @Failure	400	{object}	httputil.HTTPError
// @Failure	500	{object}	httputil.HTTPError
// @Router		/auth/current-user [get]
func (h handler) getCurrentUser(c *gin.Context) {
	userId := c.GetInt("userId")
	if userId == 0 {
		httputil.NewHTTPErrorResponse(c, http.StatusInternalServerError, "not authenticated")
	}

	user, err := h.services.Authorization.UserByID(userId)
	if err != nil {
		httputil.NewHTTPErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
