package dao

import "github.com/SXUOJ/backend/models"

// 通过问题id用户id 获取Status列表 amount是每页多少 page是页号
func GetStatusListByQid(qid string, uid string, amount string, page string) ([]*models.Result, error) {
	return nil, nil
}

// 通过提交id获取status详细
func GetStatusByAdmitId(aid string) (models.Result, error) {
	return models.Result{}, nil
}
