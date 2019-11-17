/**
 * Created by Goland.
 * User: yan.wang5<yan.wang5@transsion.com>
 * Date: 2019/11/3
 */
package http

const (
	RESP_SUCCESS      = 1000
	RESP_ERROR_SYSTEM = 1001
	RESP_ERROR_PARAM  = 1002
	RESP_ERROR_AUTH   = 1003

	RESP_ERROR_USER_NOT_FOUND = 1005
	RESP_ERROR_PASSWORD_INVALID = 1006
)

// 返回成功
func NewSuccess(data interface{}) Response {
	resp := newResponse(RESP_SUCCESS, "Success")
	resp.SetData(data)
	return resp
}

// 系统型错误
func NewSystemError(opts ...interface{}) Response {
	return newResponse(RESP_ERROR_SYSTEM, "System Error.", opts...)
}

// 参数错误
func NewParamError(opts ...interface{}) Response {
	return newResponse(RESP_ERROR_PARAM, "Params Error.", opts...)
}

// 权限错误
func NewAuthError(opts ...interface{}) Response {
	return newResponse(RESP_ERROR_AUTH, "Auth Error.", opts...)
}

// 用户不存在
func NewUserNotFoundError(opts ...interface{}) Response {
	return newResponse(RESP_ERROR_USER_NOT_FOUND, "The user is not found.", opts...)
}

// 账号密码错误
func NewPasswordInvalidError(opts ...interface{}) Response {
	return newResponse(RESP_ERROR_PASSWORD_INVALID, "The account or password is invalid.", opts...)
}
