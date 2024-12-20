package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ IpInfoModel = (*customIpInfoModel)(nil)

type (
	// IpInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIpInfoModel.
	IpInfoModel interface {
		ipInfoModel
		withSession(session sqlx.Session) IpInfoModel
	}

	customIpInfoModel struct {
		*defaultIpInfoModel
	}
)

// NewIpInfoModel returns a model for the database table.
func NewIpInfoModel(conn sqlx.SqlConn) IpInfoModel {
	return &customIpInfoModel{
		defaultIpInfoModel: newIpInfoModel(conn),
	}
}

func (m *customIpInfoModel) withSession(session sqlx.Session) IpInfoModel {
	return NewIpInfoModel(sqlx.NewSqlConnFromSession(session))
}
