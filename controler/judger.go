package controler

import (
	"github.com/SXUOJ/backend/logic"
	"github.com/SXUOJ/backend/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func CreateJudger(c *gin.Context) {
	addr := new(models.JudgerAddr)
	err := c.ShouldBindJSON(addr)
	if err != nil {
		zap.L().Error(" CreateJudger 转化失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	err = logic.CreateJudger(*addr)
	if err != nil {
		zap.L().Error("logic.CreateJudger err", zap.Error(err))
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

func GetJugerList(c *gin.Context) {
	re, err := logic.GetJudgerList()
	if err != nil {
		zap.L().Error("logic.GetJudgerList err", zap.Error(err))
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
