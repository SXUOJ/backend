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
	usergroup.GET("/get_user_info", controler.GetUserInfo)
	usergroup.PUT("/put_user_info", controler.PutUserInfo)

	//题目路由组
	quegroup := apigroup.Group("/question")
	quegroup.GET("/get/:question_id", controler.GetQuestionDetail)
	quegroup.GET("/get_list", controler.GetQuestionList)
	quegroup.POST("/submit", controler.PushQuestionJudge)

	//管理员组
	admingroup := apigroup.Group("/admin")
	admingroup.POST("/question/create", controler.CreateQuestion)
	admingroup.PUT("/question/change", controler.ChangeQuestion)
	admingroup.DELETE("/question/delete", controler.DelQuestion)

	//提交状态
	stugroup := apigroup.Group("/status")
	stugroup.GET("/get_list_by_question_id", controler.GetStatusList)
	stugroup.GET("/get_status_by_submit_id", controler.GetStatusDetail)

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
		apigroup.POST("/admin/upload/sample", func(context *gin.Context) {
			logic.SampleHandler(context.Writer, context.Request)
		})
	}
	return r
}
