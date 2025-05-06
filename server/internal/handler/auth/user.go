package auth

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/brotigen23/goph-keeper/server/internal/dto"
	"github.com/brotigen23/goph-keeper/server/internal/dto/auth/logindto"
	"github.com/brotigen23/goph-keeper/server/internal/dto/auth/registerdto"
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/service"
	"github.com/brotigen23/goph-keeper/server/pkg/auth"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.AuthService

	accessKey, refreshKey string
}

func New(
	service service.AuthService,
	accessKey, refreshKey string) *Handler {
	return &Handler{
		service:    service,
		accessKey:  accessKey,
		refreshKey: refreshKey,
	}
}

// Register 	godoc
// @Summary 	Регистрация нового пользователя
// @Description Создаёт нового пользователя и возвращает JWT токены
// @Tags 		auth
// @Accept  	json
// @Param 		input body registerdto.PostRequest true "Данные для регистрации"
// @Success 	200 {object} nil "Успешная регистрация"
// @Failure 	400 {object} string "Невалидные данные"
// @Failure 	409 {object} string "Пользователь уже существует"
// @Failure 	500 {object} string "Ошибка сервера"
// @Router 		/register 											[post]
func (h *Handler) Register(c *gin.Context) {
	response := dto.ResponseError{}
	var credentials registerdto.PostRequest
	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		response.Msg = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	user := &model.User{Login: credentials.Login, Password: credentials.Password}
	err = h.service.Register(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseError{Msg: err.Error()})
	}
	// TODO: from config
	accessExpire := time.Hour * 240
	refreshExpire := time.Hour * 72
	accessToken, refreshToken, err := auth.CreateTokens(user.ID, h.accessKey, h.refreshKey, accessExpire, refreshExpire)
	if err != nil {
		response.Msg = err.Error()
		c.JSON(http.StatusBadRequest, response)
	}

	c.Header("Authorization", "Bearer "+accessToken)
	c.Header("Refresh-Token", refreshToken)

	c.JSON(http.StatusAccepted, response)
}

// Login 	godoc
// @Summary 	Регистрация нового пользователя
// @Description Создаёт нового пользователя и возвращает JWT токены
// @Tags 		auth
// @Accept  	json
// @Param 		input body logindto.PostRequest true "Данные для регистрации"
// @Success 	200 {object} nil "Успешный вход"
// @Failure 	400 {object} error "Невалидные данные"
// @Failure 	409 {object} error "Пользователь уже существует"
// @Failure 	500 {object} error "Ошибка сервера"
// @Router 		/login	[post]
func (h *Handler) Login(c *gin.Context) {
	var credentials logindto.PostRequest
	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	user := model.User{Login: credentials.Login, Password: credentials.Password}
	err = h.service.Login(context.Background(), &user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	accessExpire := time.Hour * 240
	refreshExpire := time.Hour * 720
	log.Println(user.ID)
	accessToken, refreshToken, err := auth.CreateTokens(user.ID, h.accessKey, h.refreshKey, accessExpire, refreshExpire)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.Header("Authorization", "Bearer "+accessToken)
	c.Header("Refresh-Token", refreshToken)

	c.String(http.StatusAccepted, "All done!")

}
