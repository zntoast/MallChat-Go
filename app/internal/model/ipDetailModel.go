package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ IpDetailModel = (*customIpDetailModel)(nil)

type (
	// IpDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIpDetailModel.
	IpDetailModel interface {
		ipDetailModel
		withSession(session sqlx.Session) IpDetailModel
	}

	customIpDetailModel struct {
		*defaultIpDetailModel
	}
)

// NewIpDetailModel returns a model for the database table.
func NewIpDetailModel(conn sqlx.SqlConn) IpDetailModel {
	return &customIpDetailModel{
		defaultIpDetailModel: newIpDetailModel(conn),
	}
}

func (m *customIpDetailModel) withSession(session sqlx.Session) IpDetailModel {
	return NewIpDetailModel(sqlx.NewSqlConnFromSession(session))
}
