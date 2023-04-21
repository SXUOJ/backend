package controler

import (
	"encoding/json"
	"github.com/SXUOJ/backend/models"
	"net/http"

	"github.com/SXUOJ/backend/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	list, nums, err := logic.GetStatusList(qid, uid.(string), amount, page)
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
		"amount":      nums,
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
	var rs []models.ResultOfOneSample
	err = json.Unmarshal([]byte(re.Results), &rs)
	if err != nil {
		zap.L().Error("json.Unmarshal([]byte(re.Results),&rs) err :", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	detail := models.ResultDetail{
		SubmitID:   re.SubmitID,
		UserID:     re.UserID,
		QuestionID: re.QuestionID,
		Time:       re.Time,
		IfAC:       re.IfAC,
		CodeType:   re.CodeType,
		Source:     re.Source,
		Public:     re.Public,
		Results:    rs,
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": detail,
	})
}
