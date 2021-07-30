package errno

var (
	Ok                  = &Err{Code: 0, Message: "OK"}
	InternalServerError = &Err{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Err{Code: 10002, Message: "参数错误."}
	UserNotExits        = &Err{Code: 20001, Message: "用户不存在	"}
)
