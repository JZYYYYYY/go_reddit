package middlewares

import (
	"hellogo/web_app/controllers"
	"hellogo/web_app/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxx.xxx.xxx
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			controllers.ResponseError(ctx, controllers.CodeNeedLogin)
			ctx.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(ctx, controllers.CodeInvalidToken)
			ctx.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(ctx, controllers.CodeInvalidToken)
			ctx.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		ctx.Set(controllers.ContextUserIDKey, mc.UserID)
		ctx.Next() // 后续的处理请求的函数中，可以用过ctx.Get(ContextUserIDKey)来获取当前请求的用户信息
	}
}
