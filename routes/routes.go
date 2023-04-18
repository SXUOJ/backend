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
	r.Use(Cors())
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
	quegroup.Use(middleware.JWTAuthMiddleware())
	quegroup.GET("/get/:question_id", controler.GetQuestionDetail)
	quegroup.GET("/get_list", controler.GetQuestionList)
	quegroup.POST("/submit", controler.PushQuestionJudge)
	quegroup.GET("/search", controler.GetSearch)

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
	stugroup.Use(middleware.JWTAuthMiddleware())
	stugroup.GET("/get_list_by_question_id/:qid", controler.GetStatusList)
	stugroup.GET("/get_status_by_submit_id/:submitId", controler.GetStatusDetail)

	return r
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
