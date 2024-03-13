package entity

type ResponseEntity struct{}

type ResSucc struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResFail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewResEntity() ResponseEntity {
	return ResponseEntity{}
}

func (entity ResponseEntity) ResWithSucc(data any) ResSucc {
	res := ResSucc{
		Code:    "0000",
		Message: "Succ",
		Data:    data,
	}
	return res
}

func (entity ResponseEntity) ResWithFail(err ErrorMessage) ResFail {
	res := ResFail{
		Code:    err.Code,
		Message: err.Message,
	}
	return res
}
