package dal

import (
	"code.bigogo.com/web/goapi/pbgen/apicommon"
	"github.com/yuwenhui/goauth/config"
	"github.com/yuwenhui/goauth/models"
)

func GetTodos() ([]*apicommon.Todo, error) {
	var todos []*models.Todo
	err := config.DBConfig.Where("status = 1").Find(&todos).Error

	if err != nil {
		return nil, err
	}

	return todos, nil
}
