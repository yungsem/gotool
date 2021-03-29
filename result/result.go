package result

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

func Error(message string) Result {
	return Result{
		Code:    codeError,
		Message: message,
	}
}
