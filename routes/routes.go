package routes

import (
	"net/http"
	"web_app/controler"
	"web_app/logger"
	"web_app/logic"
	"web_app/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	})

	apigroup := r.Group("/api")
	apigroup.POST("/user/register", controler.RegisterHandler)
	apigroup.POST("/user/login", controler.LoginHandler)
	//用户路由组
	usergroup := apigroup.Group("/user")
	usergroup.Use(middleware.JWTAuthMiddleware())
	usergroup.POST("/get_user_info", controler.GetUserInfo)
	usergroup.PUT("/put_user_info", controler.PutUserInfo)

	//题目路由组
	quegroup := apigroup.Group("/question")
	quegroup.POST("/get/:question_id", controler.GetQuestionDetail)
	quegroup.POST("/get_list", controler.GetQuestionList)
	quegroup.POST("/submit", controler.PushQuestionJudge)

	// 文件上传组
	{
		apigroup.POST("/admin/upload/image/:name", func(context *gin.Context) {
			logic.Handler(context.Writer, context.Request)
		})
		apigroup.GET("/admin/upload/image/:name", func(context *gin.Context) {
			logic.Handler(context.Writer, context.Request)
		})
		apigroup.DELETE("/admin/upload/image/:imgname", func(context *gin.Context) {
			logic.Handler(context.Writer, context.Request)
		})
	}
	return r
}
