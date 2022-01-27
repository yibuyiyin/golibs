package errors

import (
	"encoding/json"
	"errors"
)

type Errors struct {
	Code    int    `json:"code" example:"4000000"`
	Message string `json:"message"`
}

var codeList = map[string]Errors{
	"internal_service_error": {5000000, "服务内部错误"},
}

func Error(key string) error {
	ret, _ := json.Marshal(codeList[key])
	return errors.New(string(ret))
}

func (e *Errors) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

func (e *Errors) SetCode(code int) {
	e.Code = code
}

func (e *Errors) SetMessage(message string) {
	e.Message = message
}

func (e *Errors) GetCode() int {
	return e.Code
}

func (e *Errors) GetMessage() string {
	return e.Message
}
