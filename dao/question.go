package dao

import "github.com/SXUOJ/backend/models"

// 通过问题id获得问题详细
func GetQuestionDetail(qid string) (*models.Question, error) {
	var questionSql models.QuestionSql
	tx := db.Where("question_id = ?", qid).First(&questionSql)
	//返回
	return &questionSql.Question, tx.Error
}

// 获取问题列表 page是页号 amount是每页数量 并且 获取每个题目是否ac
func GetQuestionList(page int, amount int) ([]*models.Question, error) {
	var (
		offset       = (page - 1) * amount
		questionSqls []models.QuestionSql
		questions    []*models.Question
	)

	db.Limit(amount).Offset(offset).Find(&questionSqls)

	for _, v := range questionSqls {
		questions = append(questions, &v.Question)
	}

	return questions, nil
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
