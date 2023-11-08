package controllers

import (
	"errors"
	"fmt"
	"hellogo/web_app/dao/mysql"
	"hellogo/web_app/logic"
	"hellogo/web_app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := ctx.ShouldBindJSON(p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("SignUp with invaild param", zap.Error(err))
		//判断err是不是alidator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMessage(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) //翻译错误
		return
	}
	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.Signup failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(ctx, CodeUserExist)
		}
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(ctx, nil)
}

// LoginHandler 处理登录请求的函数
func LoginHandler(ctx *gin.Context) {
	//1.获取请求参数和参数处理
	p := new(models.ParamLogin)
	if err := ctx.ShouldBindJSON(p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("Login with invaild param", zap.Error(err))
		//判断err是不是alidator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMessage(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) //翻译错误
		return
	}
	//2.业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(ctx, CodeUserNotExist)
			return
		}
		ResponseError(ctx, CodeInvalidPassword)
		return
	}
	//3.返回响应
	ResponseSuccess(ctx, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
