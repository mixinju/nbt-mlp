package errno

type Errno struct {
    Code    int
    Message string
}

var (
    // Common errors
    OK                  = &Errno{Code: 0, Message: "OK"}
    InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
    ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct"}
    ErrValidation       = &Errno{Code: 10003, Message: "Validation failed"}
    ErrDatabase         = &Errno{Code: 10004, Message: "Database error"}
    ErrToken            = &Errno{Code: 10005, Message: "Error occurred while signing the JSON web token"}

    // User errors
    ErrUserNotFound      = &Errno{Code: 20001, Message: "User not found"}
    ErrPasswordIncorrect = &Errno{Code: 20002, Message: "Password is incorrect"}
    ErrUserParamsInvalid = &Errno{Code: 20003, Message: "用户参数不合法"}

    // ErrFileUpload File errors
    ErrFileUpload       = &Errno{Code: 30001, Message: "File upload error"}
    ErrFileSave         = &Errno{Code: 30002, Message: "Failed to save file"}
    ErrFileOpen         = &Errno{Code: 30003, Message: "Failed to open Excel file"}
    ErrFileContentEmpty = &Errno{Code: 30004, Message: "Excel文件中有效内容为空"}

    // Token、权限相关 102 开头
    ErrTokenExpired     = &Errno{Code: 50201, Message: "Token 已过期"}
    ErrTokenSetUpFail   = &Errno{Code: 50202, Message: "Token 生成失败"}
    ErrNoToken          = &Errno{Code: 50203, Message: "No Token"}
    ErrTokenInvalid     = &Errno{Code: 50204, Message: "这不是一个Token 请重新登录"}
    ErrUnAuthorization  = &Errno{Code: 50205, Message: "没有权限执行操作"}
    ErrValidateFail     = &Errno{Code: 50206, Message: "参数校验失败"}
    ErrPasswordNotMatch = &Errno{Code: 50207, Message: "旧密码不正确"}
    ErrPasswordGenerate = &Errno{Code: 50207, Message: "生成密码错误"}
)
