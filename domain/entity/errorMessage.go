package entity

import "net/http"

type ErrorMessage struct {
	HttpCode int
	Code     string
	Message  string
}

var (
	NotFoundErr          = ErrorMessage{http.StatusNotFound, "E0001", "not found"}
	MissingFieldErr      = ErrorMessage{http.StatusBadRequest, "E0002", "missing field"}
	DbConnectErr         = ErrorMessage{http.StatusInternalServerError, "E0003", "Db op failed"}
	UserExistErr         = ErrorMessage{http.StatusPaymentRequired, "E0004", "user has existed"}
	UserNotFoundErr      = ErrorMessage{http.StatusNotFound, "E0005", "user not found"}
	PasswordIncorrectErr = ErrorMessage{http.StatusPaymentRequired, "E0006", "password incorrect"}
	TokenGenFail         = ErrorMessage{http.StatusInternalServerError, "E0007", "jwt token gen fail"}
	MissingTokenErr      = ErrorMessage{http.StatusUnauthorized, "E0008", "missing authorization token"}
	TokenInvalidErr      = ErrorMessage{http.StatusUnauthorized, "E0010", "invalid token"}
)
