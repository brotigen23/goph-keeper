package account

import (
	"context"
	"log"
	"net/http"

	dto "github.com/brotigen23/goph-keeper/server/internal/dto/account"
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.AccountService
}

func New(service service.AccountService) *Handler {
	return &Handler{
		service: service,
	}
}

// AccountsCreate 		godoc
// @Summary 	Создание нового аккаунта
// @Description Создаёт новый данные аккаунта и возвращает созданный аккаунт
// @Tags 		data
// @Security ApiKeyAuth
// @Produce  	json
// @Param 		input body dto.PostRequest true "Данные для регистрации"
// @Success 	200 {object} nil "Успешная регистрация"
// @Failure 	400 {object} string "Невалидные данные"
// @Failure 	409 {object} string "Пользователь уже существует"
// @Failure 	500 {object} string "Ошибка сервера"
// @Router 		/user/account 											[post]
func (h *Handler) Create(c *gin.Context) {
	id, exists := c.Get("userID")
	if !exists {
		c.String(http.StatusUnauthorized, "Auth error")
		return
	}

	// Приводим к нужному типу (в зависимости от того, что вы сохраняли)
	userID, ok := id.(int)
	if !ok {
		c.String(http.StatusInternalServerError, "invalid userID type")
		return
	}
	var account dto.PostRequest

	// Парсим данные аккаунта
	err := c.ShouldBindJSON(&account)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	toSave := &model.Account{
		BaseData: model.BaseData{
			UserID:   userID,
			Metadata: account.Metadata,
		},
		Login:    account.Login,
		Password: account.Password}

	// Используем userID
	log.Println(toSave)
	err = h.service.Create(context.Background(), toSave)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, toSave)
}

// GetAccount godoc
// @Summary Получить данные аккаунта
// @Description Возвращает выбранный аккаунт по ID
// @Tags data accounts
// @Security ApiKeyAuth
// @Produce  json
// @Router /user/account [get]
func (h *Handler) Get(c *gin.Context) {
}

// GetAccounts godoc
// @Summary Получить все аккаунты
// @Description Возвращает список данных аккаунтов пользователя
// @Tags data
// @Security ApiKeyAuth
// @Produce  json
// @Router /user/accounts/fetch [get]
func (h *Handler) Fetch(c *gin.Context) {
	id, exists := c.Get("userID")
	if !exists {
		c.String(http.StatusUnauthorized, "Auth error")
		return
	}

	userID, ok := id.(int)
	if !ok {
		c.String(http.StatusInternalServerError, "invalid userID type")
		return
	}

	// Используем userID
	data, err := h.service.GetUserData(context.Background(), userID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// @Security ApiKeyAuth
// @Produce  json

// UpdateAccount godoc
// @Summary 	Обновить существующий аккаунт
// @Description Возвращает список данных аккаунтов пользователя
// @Tags 		data
// @Security 	ApiKeyAuth
// @Produce  	json
// @Param 		input body dto.PutRequest true "Данные для Обновления"
// @Success 	200 {object} nil "Успешная регистрация"
// @Failure 	400 {object} string "Невалидные данные"
// @Failure 	409 {object} string "Пользователь уже существует"
// @Failure 	500 {object} string "Ошибка сервера"
// @Router 		/user/accounts/ [put]
func (h *Handler) Update(c *gin.Context) {
	id, exists := c.Get("userID")
	if !exists {
		c.String(http.StatusUnauthorized, "Auth error")
		return
	}

	// Приводим к нужному типу (в зависимости от того, что вы сохраняли)
	userID, ok := id.(int)
	if !ok {
		c.String(http.StatusInternalServerError, "invalid userID type")
		return
	}
	var account dto.PutRequest

	// Парсим данные аккаунта
	err := c.ShouldBindJSON(&account)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	toSave := &model.Account{
		Base: model.Base{
			ID: account.ID,
		},
		BaseData: model.BaseData{
			UserID:   userID,
			Metadata: account.Metadata,
		},
		Login:    account.Login,
		Password: account.Password}

	// Используем userID
	err = h.service.Update(context.Background(), toSave)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, toSave)
}
