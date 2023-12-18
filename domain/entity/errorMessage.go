package domain

import "net/http"

type ErrorMessage struct {
	HttpCode int
	Code     string
	Message  string
}

var (
	NotFoundErr     = ErrorMessage{http.StatusNotFound, "E0001", "not found"}
	MissingFieldErr = ErrorMessage{http.StatusBadRequest, "E0002", "missing field"}
	DbConnectErr    = ErrorMessage{http.StatusInternalServerError, "E0003", "Db op failed"}
	UserExistErr    = ErrorMessage{http.StatusPaymentRequired, "E0004", "user has existed"}
)
