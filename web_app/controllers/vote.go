package controllers

import (
	"hellogo/web_app/logic"
	"hellogo/web_app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//投票

func PostVoteController(ctx *gin.Context) {
	//参数校验
	p := new(models.ParamVoteData)
	if err := ctx.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) //翻译并去除掉错误提示中的结构体标识
		ResponseErrorWithMessage(ctx, CodeInvalidParam, errData)
		return
	}

	//拿到当前请求的用户ID
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ResponseError(ctx, CodeNeedLogin)
		return
	}

	//具体投票的业务逻辑
	logic.VoteForPost(userID, p)
	ResponseSuccess(ctx, nil)
}
