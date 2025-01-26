package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ItemConfigsModel = (*customItemConfigsModel)(nil)

const (
	// 物品类型 1改名卡 2徽章
	ItemConfigsTypeCard  = 1
	ItemConfigsTypeBadge = 2
)

type (
	// ItemConfigsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customItemConfigsModel.
	ItemConfigsModel interface {
		itemConfigsModel
	}

	customItemConfigsModel struct {
		*defaultItemConfigsModel
	}
)

// NewItemConfigsModel returns a model for the database table.
func NewItemConfigsModel(conn sqlx.SqlConn) ItemConfigsModel {
	return &customItemConfigsModel{
		defaultItemConfigsModel: newItemConfigsModel(conn),
	}
}
