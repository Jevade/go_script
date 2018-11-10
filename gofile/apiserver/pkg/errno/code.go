package errno

var (
	//Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occured while binding the rewuest body to the struct"}

	ErrValidation = &Errno{Code: 20001, Message: "Validation Field"}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error"}
	ErrToken      = &Errno{Code: 20003, Message: "Error occured while signing the JSON web token"}

	//User errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Err oucured while encrypting the user password"}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found"}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "Token is invalid "}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "Password is unincorrect"}
	ErrFieldEmpty        = &Errno{Code: 20105, Message: "Username or password is empty"}
)
