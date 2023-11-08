package controllers

import (
	"hellogo/web_app/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//和社区相关的

func CommunityHandler(ctx *gin.Context) {
	//查询到所有的社区（community_id,community_name）,以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy) //不轻易把服务端报错对外暴露
		return
	}
	ResponseSuccess(ctx, data)
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(ctx *gin.Context) {
	//1.获取社区id
	idStr := ctx.Param("id") //获取url参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}

	//查询到所有的社区（community_id,community_name）,以列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy) //不轻易把服务端报错对外暴露
		return
	}
	ResponseSuccess(ctx, data)
}
