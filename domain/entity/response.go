package entity

import (
	"github.com/gin-gonic/gin"
)

var Response ResponseEntity

type ResponseEntity struct{}

func (r *ResponseEntity) ResWithSucc(data any) gin.H {
	return gin.H{
		"code":    "0000",
		"message": "Succ",
		"data":    data,
	}
}

func (r *ResponseEntity) ResWithFail(err ErrorMessage) gin.H {
	return gin.H{
		"code":    err.Code,
		"message": err.Message,
	}
}
