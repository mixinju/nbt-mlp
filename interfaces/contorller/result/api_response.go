package result

import "nbt-mlp/util/errno"

type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

var OK = Response{Code: 0, Message: "success"}
var NoToken = Response{Code: errno.ErrNoToken.Code, Message: errno.ErrNoToken.Message}
var TokenExpired = Response{Code: errno.ErrTokenExpired.Code, Message: errno.ErrTokenExpired.Message}
var InvalidToken = Response{Code: errno.ErrTokenInvalid.Code, Message: errno.ErrToken.Message}
var InvalidParma = Response{Code: errno.ErrValidateFail.Code, Message: errno.ErrValidateFail.Message}
var UnAuthorized = Response{Code: errno.ErrUnAuthorization.Code, Message: errno.ErrUnAuthorization.Message}

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

type Login struct {
    UserID uint64 `json:"user_id"`
    Token  string `json:"token"`
}
