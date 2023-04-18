package dao

import (
	"fmt"

	"github.com/SXUOJ/backend/models"
)

// 通过问题id获得问题详细
func GetQuestionDetail(qid string) (*models.Question, error) {
	var questionSql models.QuestionSql
	tx := db.Where("question_id = ?", qid).First(&questionSql)
	//返回
	return &questionSql.Question, tx.Error
}

// 获取问题列表 page是页号 amount是每页数量 并且 获取每个题目是否ac
func GetQuestionList(page int, amount int, uid string) (*[]models.QueList, error) {
	var (
		offset       = (page - 1) * amount
		questionList []models.QueList
	)
	/* SELECT question_sqls.*,ac_sqls.*
		  FROM question_sqls
	 	  	LEFT JOIN ac_sqls
			  ON question_sqls.question_id = ac_sqls.question_id AND ac_sqls.user_id=uid AND ac_sqls.if_ac=1;
	*/
	result := db.Model(&models.QuestionSql{}).
		Select("question_sqls.*, ac_sqls.user_id, ac_sqls.if_ac, ac_sqls.ac_question_id").
		Joins(fmt.Sprintf("LEFT JOIN ac_sqls ON question_sqls.question_id = ac_sqls.ac_question_id AND ac_sqls.user_id='%s' AND ac_sqls.if_ac=1", uid)).
		Limit(amount).
		Offset(offset).
		Scan(&questionList)
	return &questionList, result.Error
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
