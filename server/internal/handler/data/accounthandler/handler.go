package accounthandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/brotigen23/goph-keeper/server/internal/dto"
	"github.com/brotigen23/goph-keeper/server/internal/dto/accountdto"
	"github.com/brotigen23/goph-keeper/server/internal/handler"
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
// @Summary 	Создать новый аккаунт
// @Description Создает новый аккаунт
// @Tags 		accounts
// @Security 	ApiKeyAuth
// @Produce  	json
// @Param 		input body 		accountdto.PostRequest true "Данные для сохранения"
// @Success 	200 			{object} 	nil "Успешное создание"
// @Failure 	400 			{object} 	string "Невалидные данные"
// @Failure 	401 			{object} 	string "Ошибка аутентификации"
// @Failure 	409 			{object} 	string "Конфликт создания"
// @Failure 	500 			{object} 	string "Внутренняя ошибка сервера"
// @Router 		/user/account	[post]
func (h *Handler) Post(c *gin.Context) {
	errReponse := dto.ResponseError{}
	id, exists := c.Get("userID")
	if !exists {
		errReponse.Msg = handler.ErrAuth
		c.JSON(http.StatusUnauthorized, errReponse)
		return
	}

	userID, ok := id.(int)
	if !ok {
		errReponse.Msg = handler.ErrAuth
		c.JSON(http.StatusUnauthorized, errReponse)
		return
	}
	var request accountdto.PostRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errReponse.Msg = err.Error()
		c.JSON(http.StatusUnauthorized, errReponse)
		return
	}
	account := &model.Account{
		BaseData: model.BaseData{
			UserID:   userID,
			Metadata: request.Metadata,
		},
		Login:    request.Login,
		Password: request.Password}

	err = h.service.Create(context.Background(), account)
	if err != nil {
		errReponse.Msg = err.Error()
		c.JSON(http.StatusUnauthorized, errReponse)
		return
	}
	response := accountdto.PostResponse{}
	response.Map(*account)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) Get(c *gin.Context) {
}

// GetAccounts godoc
// @Summary Получить все аккаунты
// @Description Возвращает список данных аккаунтов пользователя
// @Tags accounts
// @Security ApiKeyAuth
// @Produce  json
// @Success 	200 			{object} 	[]accountdto.GetResponse "Успешное выполнение"
// @Failure 	401 			{object} 	string "Ошибка аутентификации"
// @Failure 	500 			{object} 	string "Внутренняя ошибка сервера"
// @Router /user/accounts/fetch [get]
func (h *Handler) Fetch(c *gin.Context) {
	responseErr := dto.ResponseError{}
	id, exists := c.Get("userID")
	if !exists {
		responseErr.Msg = handler.ErrAuth
		c.JSON(http.StatusUnauthorized, responseErr)
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
		responseErr.Msg = err.Error()
		c.JSON(http.StatusUnauthorized, responseErr)
		return
	}
	response := []accountdto.GetResponse{}
	// User mapper
	for _, account := range data {
		item := accountdto.GetResponse{}
		item.Map(account)
		response = append(response, item)
	}
	c.JSON(http.StatusOK, response)
}

// UpdateAccount godoc
// @Summary 	Обновить существующий аккаунт
// @Tags 		accounts
// @Security 	ApiKeyAuth
// @Produce  	json
// @Param 		input body accountdto.PutRequest true "Данные для обновления"
// @Success 	200 			{object} 	accountdto.PutResponse "Успешное обновление"
// @Failure 	400 			{object} 	string "Невалидные данные"
// @Failure 	401 			{object} 	string "Ошибка аутентификации"
// @Failure 	500 			{object} 	string "Внутренняя ошибка сервера"
// @Router 		/user/accounts/ [put]
func (h *Handler) Put(c *gin.Context) {
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
	var request accountdto.PutRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	// TODO: mapper
	data := &model.Account{
		Base: model.Base{
			ID: request.ID,
		},
		BaseData: model.BaseData{
			UserID:   userID,
			Metadata: request.Metadata,
		},
		Login:    request.Login,
		Password: request.Password,
	}

	err = h.service.Update(context.Background(), data)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	response := accountdto.PutResponse{}
	response.Map(*data)
	c.JSON(http.StatusOK, data)
}

// Delete godoc
// @Summary 	Удалить существующий аккаунт
// @Description Удаляет аккаунт с входящим ID
// @Tags 		accounts
// @Security 	ApiKeyAuth
// @Produce  	json
// @Param 		input body accountdto.DeleteRequest true "Данные для с id записью"
// @Success 	200 			{object} 	string "Успешное удаление"
// @Failure 	400 			{object} 	string "Невалидные данные"
// @Failure 	401 			{object} 	string "Ошибка аутентификации"
// @Failure 	500 			{object} 	string "Внутренняя ошибка сервера"
// @Router 		/user/accounts/ [delete]
func (h *Handler) Delete(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
	}

	var request accountdto.DeleteRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Delete(context.Background(), userID, request.ID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "Deleted")
}

func getUserID(c *gin.Context) (int, error) {
	id, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("Auth error")
	}

	userID, ok := id.(int)
	if !ok {
		return 0, errors.New("Auth error")
	}
	return userID, nil
}
