package response

type Error struct {
	code    int
	message string
}

var (
	OK                     = Error{0, "ok"}
	SYSTEM_ERROR           = Error{100000, "系统错误"}
	PARAMS_ERROR           = Error{100001, "参数错误"}
	ACCOUNT_PASSWORD_ERROR = Error{100020, "账号密码错误"}
)
