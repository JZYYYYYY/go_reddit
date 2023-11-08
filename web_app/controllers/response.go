package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* 定义状态码
{
	"code": 10000 //程序中的错误码
	"message": "" //提示信息
	"data": ... //数据
}
*/

type ResponseData struct {
	Code    ResCode     `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseError(ctx *gin.Context, code ResCode) {
	ctx.JSON(http.StatusOK, &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	})
}

func ResponseErrorWithMessage(ctx *gin.Context, code ResCode, message interface{}) {
	ctx.JSON(http.StatusOK, &ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	})
}
