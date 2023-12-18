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
	DbConnectError  = ErrorMessage{http.StatusInternalServerError, "E0003", "Db op failed"}
)
