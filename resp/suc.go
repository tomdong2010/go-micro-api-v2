/**
 * Created by Goland.
 * User: yan.wang5<yan.wang5@transsion.com>
 * Date: 2019/11/3
 */
package resp

const RESP_SUCCESS = 1000

// 返回成功
func NewSuccess(data interface{}) Response {
	resp := newResponse(RESP_SUCCESS, "")
	resp.SetData(data)
	return resp
}
