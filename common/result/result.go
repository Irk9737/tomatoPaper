package result

import (
	"net/http"
	"tomatoPaper/web"
)

// Result 通用返回结构体
type Result struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回的数据
}

// Success 返回成功
func Success(c *web.Context, data any) {
	if data == nil {
		data = c.UserValues
		//data = gin.H{}
	}
	res := Result{}
	res.Code = int(ApiCode.SUCCESS)
	res.Message = ApiCode.GetMessage(ApiCode.SUCCESS)
	res.Data = data
	//c.JSON(http.StatusOK, res)
	_ = c.RespJSON(http.StatusOK, res)
}

// Failed 返回失败
func Failed(c *web.Context, code int, message string) {
	res := Result{}
	res.Code = code
	res.Message = message
	res.Data = c.UserValues
	//c.JSON(http.StatusOK, res)
	res.Data = c.RespData
	_ = c.RespJSON(http.StatusOK, res)
}
