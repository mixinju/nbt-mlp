package result

import "nbt-mlp/util/errno"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(data interface{}) Response {
	return Response{
		Code:    errno.OK.Code,
		Message: errno.OK.Message,
		Data:    data,
	}
}

func Error(err *errno.Errno) Response {
	return Response{
		Code:    err.Code,
		Message: err.Message,
	}
}

func ErrorWithMessage(err *errno.Errno, message string) Response {
	return Response{
		Code:    err.Code,
		Message: message,
	}
}
