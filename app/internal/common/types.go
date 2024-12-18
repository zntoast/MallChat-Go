package common

import (
	"context"
)

// ContextKey 上下文键类型
type ContextKey string

const (
	// UserIDKey 用户ID的上下文键
	UserIDKey ContextKey = "userId"
)

// GetUserIDFromContext 从上下文获取用户ID
func GetUserIDFromContext(ctx context.Context) (int64, bool) {
	value := ctx.Value(UserIDKey)
	if value == nil {
		return 0, false
	}
	userId, ok := value.(int64)
	return userId, ok
}
