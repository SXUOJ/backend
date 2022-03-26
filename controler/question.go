package controler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"web_app/logic"
	"web_app/models"
)

func GetQuestionDetail(c *gin.Context) {
	//绑定参数
	Qid := c.Param("id")
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
	code := new(models.Code)
	err := c.ShouldBindJSON(code)
	if err != nil {
		zap.L().Error("c.ShouldBindJSON(code) err..", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	jsonBack := new(bytes.Buffer)
	json.NewEncoder(jsonBack).Encode(*code)
	rsps, err := http.Post("http://101.42.237.62:9000/submit", "json", jsonBack)
	if err != nil {
		zap.L().Error("http.Post err..", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	body, err := ioutil.ReadAll(rsps.Body)
	if err != nil {
		zap.L().Error("ioutil.ReadAll err..", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err,
		})
		return
	}
	var bodyMap map[string]interface{}
	json.Unmarshal(body, &bodyMap)
	//code.CodeState = bodyMap["result"].(string)
	//err = mysql.SaveCode(code)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": 404,
	//		"msg":  err,
	//	})
	//	return
	//}
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": bodyMap["result"],
	})
	return
}
