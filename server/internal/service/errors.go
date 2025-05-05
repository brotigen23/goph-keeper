package service

import "errors"

// TODO: добавить ошибки для каждого вида данных
// избавиться от обощения
var (
	ErrNotImplement      = errors.New("Not implement")
	ErrPasswordIncorrect = errors.New("Password incorrect")
)

// TODO: сделать структуру и интерфейс для разных реализаций
func TranslateErr(err error) error {
	switch err {
	default:
		return err
	}
}
