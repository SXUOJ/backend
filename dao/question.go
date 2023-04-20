package dao

import (
	"fmt"

	"github.com/SXUOJ/backend/models"
)

// 通过问题id获得问题详细
func GetQuestionDetail(qid string) (*models.Question, error) {
	var questionSql models.QuestionSql
	res := db.Where("question_id = ?", qid).First(&questionSql)
	//返回
	return &questionSql.Question, res.Error
}

func GetSearchList(keyword string, page int, amount int, uid string) (*[]models.QueList, int64, error) {
	var (
		searchQuestionList []models.QueList
		offset             = (page - 1) * amount
		count              int64
	)

	res := db.Model(&models.QuestionSql{}).
		Select("question_sqls.*, ac_sqls.user_id, ac_sqls.if_ac, ac_sqls.ac_question_id").
		Where("question_sqls.title LIKE ?", "%"+keyword+"%").
		Limit(amount).
		Offset(offset).
		Joins(fmt.Sprintf("LEFT JOIN ac_sqls ON question_sqls.question_id = ac_sqls.ac_question_id AND ac_sqls.user_id='%s' AND ac_sqls.if_ac=1", uid)).
		Find(&searchQuestionList)
	if res.Error != nil {
		return &searchQuestionList, count, res.Error
	}

	res = db.Model(&models.QuestionSql{}).
		Where("question_sqls.title LIKE ?", "%"+keyword+"%").
		Count(&count)

	return &searchQuestionList, count, res.Error
}

// 获取问题列表 page是页号 amount是每页数量 并且 获取每个题目是否ac
func GetQuestionList(page int, amount int, uid string) (*[]models.QueList, int64, error) {
	var (
		offset       = (page - 1) * amount
		questionList []models.QueList
		count        int64
	)
	/* SELECT question_sqls.*,ac_sqls.*
		  FROM question_sqls
	 	  	LEFT JOIN ac_sqls
			  ON question_sqls.question_id = ac_sqls.question_id AND ac_sqls.user_id=uid AND ac_sqls.if_ac=1;
	*/
	res := db.Model(&models.QuestionSql{}).
		Select("question_sqls.*, ac_sqls.user_id, ac_sqls.if_ac, ac_sqls.ac_question_id").
		Joins(fmt.Sprintf("LEFT JOIN ac_sqls ON question_sqls.question_id = ac_sqls.ac_question_id AND ac_sqls.user_id='%s' AND ac_sqls.if_ac=1", uid)).
		Limit(amount).
		Offset(offset).
		Scan(&questionList)
	if res.Error != nil {
		return &questionList, count, res.Error
	}

	res = db.Model(&models.QuestionSql{}).
		Count(&count)
	return &questionList, count, res.Error
}

// 插入问题
func InsertQuestion(que models.Question) error {
	return db.Create(&models.QuestionSql{Question: que}).Error
}

// 根据问题id,以及修改后的question修改问题
func UpdateQuestion(qid string, que models.Question) error {
	return db.Model(&models.QuestionSql{}).Where("question_id = ?", qid).Updates(models.QuestionSql{Question: que}).Error
}

// 根据问题id删除问题
func DeleteQuestion(qid string) error {
	return db.Where("question_id = ?", qid).Unscoped().Delete(&models.QuestionSql{}).Error
}

// 插入AC表
func InsertAc(ac models.Ac) (err error) {
	var ac1 models.AcSql
	if err = db.Where("question_id = ?", ac.AcQuestionID).First(&ac1).Error; err != nil {
		db.Create(&models.AcSql{Ac: ac})
	}
	return err
}
