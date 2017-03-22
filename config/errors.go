package config

import "net/http"

type ApiError struct {
	Code        int    `json:"code"`
	HttpCode    int    `json:"http"`
	Message     string `json:"message"`
	Info        string `json:"info"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(err error) ApiError {
	return ApiError{500, http.StatusInternalServerError, err.Error(), ""}
}

var ErrNoUser = ApiError{120, http.StatusInternalServerError, "Не указан пользователь", ""}
var ErrCreateUser = ApiError{121, http.StatusInternalServerError, "Ошибка создания пользователя", ""}

var ErrGameIdWrong = ApiError{130, http.StatusBadRequest, "Нет игры с данным ID", ""}
var ErrGameTitleWrong = ApiError{131, http.StatusBadRequest, "Неверный заголовок игры", ""}
var ErrCreateGame = ApiError{132, http.StatusBadRequest, "Ошибка создания игры", ""}