package auth

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/brotigen23/goph-keeper/server/internal/dto/auth/login"
	"github.com/brotigen23/goph-keeper/server/internal/dto/auth/register"
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
// @Param 		input body register.Request true "Данные для регистрации"
// @Success 	200 {object} nil "Успешная регистрация"
// @Failure 	400 {object} string "Невалидные данные"
// @Failure 	409 {object} string "Пользователь уже существует"
// @Failure 	500 {object} string "Ошибка сервера"
// @Router 		/register 											[post]
func (h *Handler) Register(c *gin.Context) {
	var credentials register.Request
	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	user := &model.User{Login: credentials.Login, Password: credentials.Password}
	err = h.service.Register(context.Background(), user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	accessExpire := time.Hour * 24
	refreshExpire := time.Hour * 72
	accessToken, refreshToken, err := auth.CreateTokens(user.ID, h.accessKey, h.refreshKey, accessExpire, refreshExpire)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.Header("Authorization", "Bearer "+accessToken)
	c.Header("Refresh-Token", refreshToken)

	c.String(http.StatusAccepted, "All done!")
}

// Login 	godoc
// @Summary 	Регистрация нового пользователя
// @Description Создаёт нового пользователя и возвращает JWT токены
// @Tags 		auth
// @Accept  	json
// @Param 		input body login.Request true "Данные для регистрации"
// @Success 	200 {object} nil "Успешная регистрация"
// @Failure 	400 {object} string "Невалидные данные"
// @Failure 	409 {object} string "Пользователь уже существует"
// @Failure 	500 {object} string "Ошибка сервера"
// @Router 		/login	[post]
func (h *Handler) Login(c *gin.Context) {
	var credentials login.Request
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
