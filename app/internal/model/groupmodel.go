package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupModel = (*customGroupModel)(nil)

type (
	// GroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupModel.
	GroupModel interface {
		groupModel
		GetMembers(groupId int64) ([]int64, error)
		AddMember(groupId, userId int64) error
		RemoveMember(groupId, userId int64) error
		FindOneByGroupIdUserId(ctx context.Context, groupId int64, userId int64) (*GroupMember, error)
	}

	customGroupModel struct {
		*defaultGroupModel
	}
)

// NewGroupModel returns a model for the database table.
func NewGroupModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupModel {
	return &customGroupModel{
		defaultGroupModel: newGroupModel(conn, c, opts...),
	}
}

// GetMembers 获取群组成员列表
func (m *customGroupModel) GetMembers(groupId int64) ([]int64, error) {
	var members []int64
	query := `select user_id from group_member where group_id = ?`
	err := m.CachedConn.QueryRowsNoCacheCtx(context.Background(), &members, query, groupId)
	return members, err
}

// AddMember 添加群组成员
func (m *customGroupModel) AddMember(groupId, userId int64) error {
	query := `insert into group_member (group_id, user_id, join_time) values (?, ?, ?)`
	_, err := m.CachedConn.ExecCtx(context.Background(), func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, groupId, userId, time.Now().Unix())
	})
	return err
}

// RemoveMember 移除群组成员
func (m *customGroupModel) RemoveMember(groupId, userId int64) error {
	query := `delete from group_member where group_id = ? and user_id = ?`
	_, err := m.CachedConn.ExecCtx(context.Background(), func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, groupId, userId)
	})
	return err
}

func (m *customGroupModel) FindOneByGroupIdUserId(ctx context.Context, groupId int64, userId int64) (*GroupMember, error) {
	var member GroupMember
	query := `select * from group_member where group_id = ? and user_id = ?`
	err := m.CachedConn.QueryRowNoCacheCtx(ctx, &member, query, groupId, userId)
	if err != nil {
		return nil, err
	}
	return &member, nil
}
