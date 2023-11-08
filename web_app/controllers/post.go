package controllers

import (
	"hellogo/web_app/logic"
	"hellogo/web_app/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePostHandler 创建帖子的处理函数
func CreatePostHandler(ctx *gin.Context) {
	//1.参数获取和校验
	p := new(models.Post)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Debug("ctx.ShouldBindJSON error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	//从ctx取到当前发请求的用户ID
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ResponseError(ctx, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(ctx, nil)
}

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(ctx *gin.Context) {
	//1.获取参数(从url获取帖子id)
	pidStr := ctx.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	//2.根据id查数据库
	data, err := logic.GetPostDetailByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostDetailByID failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(ctx, data)
}

// GetPostListHandler 获取帖子列表的处理函数
func GetPostListHandler(ctx *gin.Context) {
	//获取分页参数
	page, size := getPageInfo(ctx)

	//获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}

	//返回响应
	ResponseSuccess(ctx, data)
}

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]

// GetPostListHandler2 获取帖子列表的处理函数(根据前端参数动态获取帖子列表)
// 1.获取参数
// 2.去redis查询id列表
// 3.根据id去数据库查询帖子详细信息
func GetPostListHandler2(ctx *gin.Context) {
	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
	// 初始化结构体时指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, // magic string
	}
	//c.ShouldBind()  根据请求的数据类型选择相应的方法去获取数据
	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := ctx.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostListNew(p) // 更新：合二为一
	// 获取数据
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, data)
	// 返回响应
}

/*
// GetCommunityPostListHandler 根据社区去查询帖子列表
func GetCommunityPostListHandler(ctx *gin.Context) {
	// 初始化结构体时指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, // magic string
	}
	//c.ShouldBind()  根据请求的数据类型选择相应的方法去获取数据
	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := ctx.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityPostList(p)
	// 获取数据
	if err != nil {
		zap.L().Error("logic.GetCommunityPostList failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, data)
	// 返回响应
}
*/
