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

	// File errors
	ErrFileUpload       = &Errno{Code: 30001, Message: "File upload error"}
	ErrFileSave         = &Errno{Code: 30002, Message: "Failed to save file"}
	ErrFileOpen         = &Errno{Code: 30003, Message: "Failed to open Excel file"}
	ErrFileContentEmpty = &Errno{Code: 3004, Message: "Excel文件中有效内容为空"}
)
