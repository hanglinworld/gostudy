package response

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

// 成功结构
func Success(data interface{}) Response {

	return Response{Data: data}
}

// 失败
func Fail(error Error, errors interface{}) Response {
	code, message := error.code, error.message

	return Response{Code: code, Message: message, Errors: errors}
}

// 对 success 的处理
// success 总是会返回 code 为 0，message  和 errors 为 null 的数据，其为默认值

// 对 fail 的处理，可能存在下边的格式
// errors 可能存在下边的情况
// errors := map[string]string{"filed1":"field1不能为空", "field2":"field2 长度不能超过 32 位"}  在表单提交的时候，对应到具体的 input 框的时候，比较好用
// errors := []string{"账号密码错误"} 在请求 api 的时候，不需要对应 input 框这样的场景
// errors := nil 不需要展示到页面里
// message message 不为 nil 时候，可以作为 message 提示信息飘一下