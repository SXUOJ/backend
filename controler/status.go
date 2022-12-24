package controler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"web_app/logic"
)

func GetStatusList(c *gin.Context) {
	uid, ok := c.Get("user_id")
	if !ok {
		zap.L().Error("can't find userid")
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "can't find userid",
		})
		return
	}
	qid := c.Param("qid")
	page := c.Query("page")
	amount := c.Query("amount")
	list, err := logic.GetStatusList(qid, uid.(string), amount, page)
	if err != nil {
		zap.L().Error("logic.GetStatusList err :", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        200,
		"result_list": list,
	})
}

func GetStatusDetail(c *gin.Context) {
	sid := c.Param("submitId")
	re, err := logic.GetStatusDetail(sid)
	if err != nil {
		zap.L().Error("logic.GetStatusDetail err :", zap.Error(err))
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
