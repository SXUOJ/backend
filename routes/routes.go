package routes

import (
	"net/http"
	"web_app/controler"
	"web_app/logger"
	"web_app/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	})

	apigroup := r.Group("api")
	apigroup.POST("/register", controler.RegisterHandler)
	apigroup.POST("/login", controler.LoginHandler)
	//用户路由组
	usergroup := apigroup.Group("user")
	usergroup.Use(middleware.JWTAuthMiddleware())
	usergroup.POST("/get_user_info", controler.GetUserInfo)
	usergroup.POST("/put_user_info", controler.PutUserInfo)

	r.Run()
	return r
}
