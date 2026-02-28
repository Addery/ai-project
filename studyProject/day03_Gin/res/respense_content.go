package res

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
type Code int

const (
	DefaultCode    Code = 0
	RoleErrCode    Code = 1001
	NetWordErrCode Code = 1002
)

var CodeMap = map[Code]string{
	DefaultCode:    "默认",
	RoleErrCode:    "权限错误",
	NetWordErrCode: "网络错误",
}

func init() {

}

func response(c *gin.Context, code Code, data any, msg string) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Ok(c *gin.Context, code Code, data any, msg string) {
	response(c, code, data, msg)
}

func OkWithMsg(c *gin.Context, msg string) {
	response(c, DefaultCode, gin.H{}, msg)
}

func OkWithData(c *gin.Context, data any) {
	response(c, DefaultCode, data, "")
}

func Fail(c *gin.Context, code Code, data any, msg string) {
	response(c, code, data, msg)
}

func FailWithMsg(c *gin.Context, msg string) {
	response(c, DefaultCode, gin.H{}, msg)
}

func FailWithCode(c *gin.Context, code Code) {
	response(c, code, gin.H{}, "")
}
