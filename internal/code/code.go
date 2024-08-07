package code

import (
	_ "embed"

	"go-gin-api-simple/configs"
)

//go:embed code.go
var ByteCodeFile []byte

// Failure 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 描述信息
}

const (
	ServerError        = 10101
	TooManyRequests    = 10102
	ParamBindError     = 10103
	AuthorizationError = 10104
	UrlSignError       = 10105
	CacheSetError      = 10106
	CacheGetError      = 10107
	CacheDelError      = 10108
	CacheNotExist      = 10109
	ResubmitError      = 10110
	HashIdsEncodeError = 10111
	HashIdsDecodeError = 10112
	RBACError          = 10113
	RedisConnectError  = 10114
	MySQLConnectError  = 10115
	WriteConfigError   = 10116
	SendEmailError     = 10117
	MySQLExecError     = 10118
	GoVersionError     = 10119
	SocketConnectError = 10120
	SocketSendError    = 10121

	UserResetPasswordError      = 20205
	UserLoginError              = 20206
	UserSignupError             = 20204
	UserLogOutError             = 20207
	UserModifyPasswordError     = 20208
	UserOldPasswordError        = 20210
	UserModifyPersonalInfoError = 20209
	UserDetailError             = 20213

	PrizeCreateError = 20501
	PrizeListError   = 20502
	PrizeDeleteError = 20503
	PrizeUpdateError = 20504
	PrizeDetailError = 20505
	PrizeModifyError = 20506
)

func Text(code int) string {
	lang := configs.Get().Language.Local

	if lang == configs.ZhCN {
		return zhCNText[code]
	}

	if lang == configs.EnUS {
		return enUSText[code]
	}

	return zhCNText[code]
}
