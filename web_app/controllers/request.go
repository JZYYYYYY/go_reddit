package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUser 获取当前登录的用户ID
func getCurrentUserID(ctx *gin.Context) (userID int64, err error) {
	uid, ok := ctx.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func getPageInfo(ctx *gin.Context) (int64, int64) {
	pageStr := ctx.Query("page")
	sizeStr := ctx.Query("size")

	var (
		size int64
		page int64
		err  error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
