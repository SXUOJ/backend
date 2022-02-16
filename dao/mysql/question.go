package mysql

import "web_app/models"

func GetQuestionDetail(Qid string) (*models.Question, error) {
	//查库
	//sql语句
	sqlStr := "select title from question where id=?"
	sqlStr2 := "select description, input, output ,sampleInput ,source from question where id=?"
	sqlStr3 := "select id, timeLimit, memLimit ,ioMode ,createBy ,level ,tags from question where id=?"
	sqlStr4 := "select ac, wa from question where id=?"
	//数据分类
	title := ""
	context := new(models.Context)
	info := new(models.Information)
	static := new(models.Statistic)
	//查询
	err := db.Get(&title, sqlStr, Qid)
	err = db.Get(context, sqlStr2, Qid)
	err = db.Get(info, sqlStr3, Qid)
	err = db.Get(static, sqlStr4, Qid)
	if err != nil {
		return nil, err
	}
	//拼接数据
	que := new(models.Question)
	que.Title = title
	que.Context = context
	que.Information = info
	que.Statistic = static
	//返回
	return que, nil
}

func GetQuestionList(page int, amount int) ([]*models.UserInMysql, error) {
	sqlStr := `select
	id, title, tags
	from question
    ORDER BY id
	limit ?,?
	`
	var data []*models.UserInMysql
	err := db.Select(&data, sqlStr, (page-1)*amount, amount)
	if err != nil {
		return nil, err
	}
	return data, nil
}
