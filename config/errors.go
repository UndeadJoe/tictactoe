package config

import "net/http"

type ApiError struct {
	Code        int    `json:"errorCode"`
	HttpCode    int    `json:"-"`
	Message     string `json:"errorMsg"`
	Info        string `json:"errorInfo"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(err error) *ApiError {
	return &ApiError{0, http.StatusInternalServerError, err.Error(), ""}
}

var ErrGameIdWrong = &ApiError{130, http.StatusBadRequest, "Wrong Game Id", ""}
var ErrGameTitleWrong = &ApiError{131, http.StatusBadRequest, "Wrong Game title", ""}