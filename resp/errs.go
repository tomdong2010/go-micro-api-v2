/**
 * Created by Goland.
 * User: yan.wang5<yan.wang5@transsion.com>
 * Date: 2019/11/3
 */
package resp

const (
	RESP_ERROR_SYSTEM = 1001
	RESP_ERROR_PARAM  = 1002
	RESP_ERROR_AUTH   = 1003
)

// 系统型错误
func NewSystemError(msg string, opts ...interface{}) Response {
	return newResponse(RESP_ERROR_SYSTEM, msg, opts...)
}

// 参数错误
func NewParamError(msg string, opts ...interface{}) Response {
	return newResponse(RESP_ERROR_PARAM, msg, opts...)
}

// 权限错误
func NewAuthError(msg string, opts ...interface{}) Response {
	return newResponse(RESP_ERROR_AUTH, msg, opts...)
}
