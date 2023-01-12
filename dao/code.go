package dao

import "github.com/SXUOJ/backend/models"

// InsertSolution 插入提交结果
func InsertSolution(Solution models.Solution) error {
	return db.Create(&models.SolutionSql{Solution: Solution}).Error
}
