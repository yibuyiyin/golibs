package errors

import (
	"encoding/json"
	"errors"
)

type Errors struct {
	ErrCode int    `json:"errCode" example:"4000000"`
	Message string `json:"message"`
}

var errCodeList = map[string]Errors{
	"internal_service_error": {5000000, "服务内部错误"},
}

func Error(key string) error {
	ret, _ := json.Marshal(errCodeList[key])
	return errors.New(string(ret))
}

func (e *Errors) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

func (e *Errors) SetErrCode(errCode int) {
	e.ErrCode = errCode
}

func (e *Errors) SetMessage(message string) {
	e.Message = message
}

func (e *Errors) GetErrCode() int {
	return e.ErrCode
}

func (e *Errors) GetMessage() string {
	return e.Message
}
