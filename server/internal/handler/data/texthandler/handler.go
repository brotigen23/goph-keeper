package texthandler

import (
	"context"
	"log"
	"net/http"

	"github.com/brotigen23/goph-keeper/server/internal/dto"
	"github.com/brotigen23/goph-keeper/server/internal/dto/accountdto"
	"github.com/brotigen23/goph-keeper/server/internal/dto/textdto"
	"github.com/brotigen23/goph-keeper/server/internal/handler"
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.TextService
}

func New(service service.TextService) *Handler {
	return &Handler{
		service: service,
	}
}

// TextCreate 		godoc
// @Summary 	Создать новую запись
// @Tags 		text
// @Security 	ApiKeyAuth
// @Produce  	json
// @Param 		input body 		textdto.PostRequest true "Данные для сохранения"
// @Success 	200 			{object} 	nil "Успешное создание"
// @Failure 	400 			{object} 	string "Невалидные данные"
// @Failure 	401 			{object} 	string "Ошибка аутентификации"
// @Failure 	409 			{object} 	string "Конфликт создания"
// @Failure 	500 			{object} 	string "Внутренняя ошибка сервера"
// @Router 		/user/text	[post]
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
	var data textdto.PostRequest

	err := c.ShouldBindJSON(&data)
	if err != nil {
		errReponse.Msg = err.Error()
		c.JSON(http.StatusUnauthorized, errReponse)
		return
	}
	toSave := &model.TextData{
		BaseData: model.BaseData{
			UserID:   userID,
			Metadata: data.Metadata,
		},
		Data: data.Data,
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

// GetText godoc
// @Summary Получить все данные
// @Tags text
// @Security ApiKeyAuth
// @Produce  json
// @Success 	200 			{object} 	nil "Успешное выполнение"
// @Failure 	401 			{object} 	string "Ошибка аутентификации"
// @Failure 	500 			{object} 	string "Внутренняя ошибка сервера"
// @Router /user/text/fetch [get]
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

// UpdateText godoc
// @Summary 	Обновить существующую запись
// @Tags 		text
// @Security 	ApiKeyAuth
// @Produce  	json
// @Param 		input body textdto.PutRequest true "Данные для обновления"
// @Success 	200 			{object} 	nil "Успешное обновление"
// @Failure 	400 			{object} 	string "Невалидные данные"
// @Failure 	401 			{object} 	string "Ошибка аутентификации"
// @Failure 	500 			{object} 	string "Внутренняя ошибка сервера"
// @Router 		/user/text/ [put]
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
	var data textdto.PutRequest

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	toSave := &model.TextData{
		Base: model.Base{
			ID: data.ID,
		},
		BaseData: model.BaseData{
			UserID:   userID,
			Metadata: data.Metadata,
		},
		Data: data.Data,
	}

	err = h.service.Update(context.Background(), toSave)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, toSave)
}

// DeleteText godoc
// @Summary 	Удалить существующую запись
// @Tags 		text
// @Security 	ApiKeyAuth
// @Produce  	json
// @Param 		input body textdto.DeleteRequest true "Данные для обновления"
// @Success 	200 			{object} 	nil "Успешное удаление"
// @Failure 	400 			{object} 	string "Невалидные данные"
// @Failure 	401 			{object} 	string "Ошибка аутентификации"
// @Failure 	500 			{object} 	string "Внутренняя ошибка сервера"
// @Router 		/user/text/ [delete]
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
