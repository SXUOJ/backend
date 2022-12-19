package mysql

import "web_app/models"

// 通过问题id获得问题详细
func GetQuestionDetail(Qid string) (*models.Question, error) {
	//查库
	//sql语句
	sqlStr := "select title from question where id=?"
	sqlStr2 := "select description, input, output,sampleOutput ,sampleInput ,source from question where id=?"
	sqlStr3 := "select id, timeLimit, memLimit ,ioMode ,createBy ,level ,tags from question where id=?"
	sqlStr4 := "select ac, wa from question where id=?"
	//数据分类
	title := ""
	context := new(models.Context)
	info := new(models.Information)
	limit := new(models.Limit)
	//查询
	err := db.Get(&title, sqlStr, Qid)
	err = db.Get(context, sqlStr2, Qid)
	err = db.Get(info, sqlStr3, Qid)
	err = db.Get(limit, sqlStr4, Qid)
	if err != nil {
		return nil, err
	}
	//拼接数据
	que := new(models.Question)
	que.Title = title
	que.Context = *context
	que.Information = *info
	que.Limit = *limit
	//返回
	return que, nil
}

// 获取问题列表 page是页号 amount是每页数量 并且 获取每个题目是否ac
func GetQuestionList(page int, amount int) ([]*models.Question, error) {
	sqlStr := `select
	id, title, tags, que_id
	from question
    ORDER BY que_id
	limit ?,?
	`
	var data []*models.Question
	err := db.Select(&data, sqlStr, (page-1)*amount, amount)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 插入问题
func InsertQuestion(que models.Question) error {
	return nil
}

// 根据问题id修改问题
func UpdateQuestion(qid string) error {
	return nil
}

// 根据问题id删除问题
func DeleteQuestion(qid string) error {
	return nil
}
