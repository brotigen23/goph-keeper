package cardhandler

import (
	"context"
	"log"
	"net/http"

	"github.com/brotigen23/goph-keeper/server/internal/dto"
	"github.com/brotigen23/goph-keeper/server/internal/dto/accountdto"
	"github.com/brotigen23/goph-keeper/server/internal/dto/cardsdto"
	"github.com/brotigen23/goph-keeper/server/internal/handler"
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.CardService
}

func New(service service.CardService) *Handler {
	return &Handler{
		service: service,
	}
}

// CardCreate 		godoc
// @Summary 	Создать данные банковской карты
// @Tags 		cards
// @Security 	ApiKeyAuth
// @Produce  	json
// @Param 		input body 		cardsdto.PostRequest true "Данные для сохранения"
// @Success 	200 			{object} 	nil "Успешное создание"
// @Failure 	400 			{object} 	string "Невалидные данные"
// @Failure 	401 			{object} 	string "Ошибка аутентификации"
// @Failure 	409 			{object} 	string "Конфликт создания"
// @Failure 	500 			{object} 	string "Внутренняя ошибка сервера"
// @Router 		/user/cards	[post]
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
	var data cardsdto.PostRequest

	err := c.ShouldBindJSON(&data)
	if err != nil {
		errReponse.Msg = err.Error()
		c.JSON(http.StatusUnauthorized, errReponse)
		return
	}
	toSave := &model.CardData{
		BaseData: model.BaseData{
			UserID:   userID,
			Metadata: data.Metadata,
		},
		Number:         data.Number,
		CardholderName: data.CardholderName,
		Expire:         data.Exipre,
		CVV:            data.CVV,
	}
	err = h.service.Create(context.Background(), toSave)
	if err != nil {
		errReponse.Msg = err.Error()
		c.JSON(http.StatusUnauthorized, errReponse)
		return
	}

	c.JSON(http.StatusOK, toSave)
}

func (h *Handler) Get(c *gin.Context) {
}

// GetCards godoc
// @Summary Получить все бинарные данные
// @Tags cards
// @Security ApiKeyAuth
// @Produce  json
// @Success 	200 			{object} 	nil "Успешное выполнение"
// @Failure 	401 			{object} 	string "Ошибка аутентификации"
// @Failure 	500 			{object} 	string "Внутренняя ошибка сервера"
// @Router /user/cards/fetch [get]
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

	c.JSON(http.StatusOK, data)
}

// UpdateCard godoc
// @Summary 	Обновить существующую карту
// @Tags 		cards
// @Security 	ApiKeyAuth
// @Produce  	json
// @Param 		input body cardsdto.PutRequest true "Данные для Обновления"
// @Success 	200 			{object} 	nil "Успешное обновление"
// @Failure 	400 			{object} 	string "Невалидные данные"
// @Failure 	401 			{object} 	string "Ошибка аутентификации"
// @Failure 	500 			{object} 	string "Внутренняя ошибка сервера"
// @Router 		/user/cards/ [put]
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
	var data cardsdto.PutRequest

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	toSave := &model.CardData{
		Base: model.Base{
			ID: data.ID,
		},
		BaseData: model.BaseData{
			UserID:   userID,
			Metadata: data.Metadata,
		},
		Number:         data.Number,
		CardholderName: data.CardholderName,
		Expire:         data.Exipre,
		CVV:            data.CVV,
	}

	err = h.service.Update(context.Background(), toSave)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, toSave)
}

// DeleteCard godoc
// @Summary 	Удалить существующую запись
// @Description Удаляет запись с входящим ID
// @Tags 		cards
// @Security 	ApiKeyAuth
// @Produce  	json
// @Param 		input body cardsdto.DeleteRequest true "Данные с id записью"
// @Success 	200 			{object} 	nil "Успешное удаление"
// @Failure 	400 			{object} 	string "Невалидные данные"
// @Failure 	401 			{object} 	string "Ошибка аутентификации"
// @Failure 	500 			{object} 	string "Внутренняя ошибка сервера"
// @Router 		/user/cards/ [delete]
func (h *Handler) Delete(c *gin.Context) {
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
	var request accountdto.DeleteRequest
	log.Println(userID)
	err := c.ShouldBindJSON(&request)
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
