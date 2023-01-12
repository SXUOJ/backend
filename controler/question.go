package controler

import (
	"net/http"
	"strconv"

	"github.com/SXUOJ/backend/logic"
	"github.com/SXUOJ/backend/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	uid, ok := c.Get("user_id")
	if !ok {
		zap.L().Error(" GetQuestionList 转化失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  zap.Error(err),
		})
		return
	}
	if err != nil {
		zap.L().Error(" GetQuestionList 转化失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  zap.Error(err),
		})
		return
	}
	//逻辑层处理
	data, err := logic.GetQuestionList(Page, Amount, uid.(string))
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
	var code models.Code
	err := c.ShouldBindJSON(&code)
	if err != nil {
		zap.L().Error(" PushQuestionJudge 转化失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	re, err := logic.PushJudge(code)
	if err != nil {
		zap.L().Error("PushJudge 失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": re,
	})
}

func CreateQuestion(c *gin.Context) {
	que := new(models.Question)
	err := c.ShouldBindJSON(que)
	if err != nil {
		zap.L().Error(" CreateQuestion 转化失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	err = logic.CreateQuestion(*que)
	if err != nil {
		zap.L().Error("logic.CreateQuestion err", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func ChangeQuestion(c *gin.Context) {
	que := new(models.Question)
	err := c.ShouldBindJSON(que)
	if err != nil {
		zap.L().Error("changeQuestion 转化失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	err = logic.ChangeQuestion(que.Information.QuestionID, *que)
	if err != nil {
		zap.L().Error("logic.ChangeQuestion err：", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func DelQuestion(c *gin.Context) {
	qid := c.Param("id")
	err := logic.DelQuestion(qid)
	if err != nil {
		zap.L().Error("logic.DelQuestion err：", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
