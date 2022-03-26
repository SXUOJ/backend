package mysql

import (
	"go.uber.org/zap"
	"web_app/models"
)

func SaveCode(code *models.Code) error {
	sqlStr := `insert into code (codeQId,codeUserId,codeType,codeSource,codeState) values(?,?,?,?,?)`
	_, err := db.Exec(sqlStr, code.CodeQId, code.CodeUserId, code.CodeType, code.CodeSource, code.CodeState)
	if err != nil {
		zap.L().Error("SaveCode err", zap.Error(err))
		return err
	}
	return nil
}
