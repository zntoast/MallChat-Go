package model

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

var ErrUserExists = fmt.Errorf("用户已存在")
var ErrInvalidPassword = fmt.Errorf("密码错误")
var ErrInvalidCode = fmt.Errorf("验证码错误或已过期")
