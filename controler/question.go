package controler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"web_app/logic"
	"web_app/models"
)

func GetQuestionDetail(c *gin.Context) {
	//绑定参数
	Qid := c.Param("question_id")
	//业务
	que := new(models.Question)
	var err error
	que, err = logic.GetQuestionDetail(Qid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"code":          202,
		"msg":           "ok",
		"question_info": que,
	})
	return
}

func GetQuestionList(c *gin.Context) {
	//绑定参数
	page := c.Query("page")
	amount := c.Query("amount")
	Page, err := strconv.Atoi(page)
	Amount, err := strconv.Atoi(amount)
	if err != nil {
		zap.L().Error(" GetQuestionList 转化失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  zap.Error(err),
		})
		return
	}
	//逻辑层处理
	data, err := logic.GetQuestionList(Page, Amount)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"code":          202,
		"msg":           "ok",
		"question_list": data,
	})
	return
}

func PushQuestionJudge(c *gin.Context) {

}
