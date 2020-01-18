package result

import (
	"github.com/json-iterator/go"
)

const (
	codeSuccess    = 0
	codeError      = -1
	messageSuccess = "成功"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *Result) toJSON() string {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(r)

	if err != nil {
		return ErrorJSON(err.Error())
	}
	return string(b)
}

func Success(data interface{}) Result {
	r := Result{
		Code:    codeSuccess,
		Message: messageSuccess,
	}
	if data != nil {
		r.Data = data
	}

	return r
}

func SuccessJSON(data interface{}) string {
	r := Success(data)
	return r.toJSON()
}

func Error(message string) Result {
	return Result{
		Code:    codeError,
		Message: message,
	}
}

func ErrorJSON(message string) string {
	r := Error(message)
	return r.toJSON()
}
