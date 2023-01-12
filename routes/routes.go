package routes

import (
	"net/http"

	"github.com/SXUOJ/backend/controler"
	"github.com/SXUOJ/backend/logger"
	"github.com/SXUOJ/backend/logic"
	"github.com/SXUOJ/backend/middleware"
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
	apigroup.Use(middleware.JWTAuthMiddleware())
	quegroup.GET("/get/:question_id", controler.GetQuestionDetail)
	quegroup.GET("/get_list", controler.GetQuestionList)
	quegroup.POST("/submit", controler.PushQuestionJudge)

	//管理员组
	admingroup := apigroup.Group("/admin")
	admingroup.Use(middleware.JWTAuthMiddleware())
	admingroup.POST("/question/create", controler.CreateQuestion)
	admingroup.PUT("/question/change", controler.ChangeQuestion)
	admingroup.DELETE("/question/delete/:id", controler.DelQuestion)
	{
		//判题机组
		judger := admingroup.Group("/judger")
		judger.POST("/create", controler.CreateJudger)
		judger.POST("/get_list", controler.GetJugerList)
	}
	// 文件上传组
	{
		admingroup.POST("/upload/image/:name", func(context *gin.Context) {
			logic.Handler(context.Writer, context.Request)
		})
		apigroup.GET("/admin/upload/image/:name", func(context *gin.Context) {
			logic.Handler(context.Writer, context.Request)
		})
		admingroup.DELETE("/upload/image/:imgname", func(context *gin.Context) {
			logic.Handler(context.Writer, context.Request)
		})
		admingroup.POST("/upload/sample/:name", func(context *gin.Context) {
			logic.SampleHandler(context.Writer, context.Request)

		})
		apigroup.GET("/admin/upload/sample/:name", func(context *gin.Context) {
			logic.SampleHandler(context.Writer, context.Request)
		})
		admingroup.DELETE("/upload/sample/:name", func(context *gin.Context) {
			logic.SampleHandler(context.Writer, context.Request)
		})
	}

	//提交状态
	stugroup := apigroup.Group("/status")
	apigroup.Use(middleware.JWTAuthMiddleware())
	stugroup.GET("/get_list_by_question_id/:qid", controler.GetStatusList)
	stugroup.GET("/get_status_by_submit_id/:submitId", controler.GetStatusDetail)

	return r
}
