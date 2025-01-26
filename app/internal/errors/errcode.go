package errors

//错误码定义
//错误码格式：前3位http状态码，后6位业务错误码
type ErrorCode uint32

const (
	ErrAUTHFAILED       ErrorCode = 401000001 // 鉴权失败
	ErrStatusBadRequest ErrorCode = 400000002 // 参数有误
	ErrDataNotFound     ErrorCode = 404000001 // 数据不存在
)

const (
	SysInternalError ErrorCode = 500000001 // 系统内部错误
	SysDBError       ErrorCode = 500000003 // 数据库错误
	SysReadFileError ErrorCode = 500000006 // 文件错误
	SysOtherError    ErrorCode = 500000007 // 其他错误
)

const (
// 业务错误码
)
