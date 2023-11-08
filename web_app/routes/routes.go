package routes

import (
	"hellogo/web_app/controllers"
	"hellogo/web_app/logger"
	"hellogo/web_app/middlewares"
	"hellogo/web_app/settings"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2, 1))

	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})

	v1 := r.Group("/api/v1")

	//注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)

	//登录业务路由
	v1.POST("/login", controllers.LoginHandler)

	v1.GET("/post/:id", controllers.GetPostDetailHandler)
	v1.GET("/posts/", controllers.GetPostListHandler)
	v1.GET("/posts2/", controllers.GetPostListHandler2)
	v1.GET("/community", controllers.CommunityHandler)
	v1.GET("/community/:id", controllers.CommunityDetailHandler)

	//应用jwt中间件
	v1.Use(middlewares.JWTAuthMiddleware())

	//发布帖子
	v1.POST("/post", controllers.CreatePostHandler)
	// 投票
	v1.POST("/vote", controllers.PostVoteController)

	pprof.Register(r) //注册pprof相关路由

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "404",
		})
	})

	return r
}
