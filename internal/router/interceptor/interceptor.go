package interceptor

import (
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/proposal"
	"go-gin-api-simple/internal/repository/mysql"
	"go-gin-api-simple/internal/repository/redis"
	"go-gin-api-simple/internal/services/user"

	"go.uber.org/zap"
)

var _ Interceptor = (*interceptor)(nil)

type Interceptor interface {
	// CheckLogin 验证是否登录
	CheckLogin(ctx core.Context) (info proposal.SessionUserInfo, err core.BusinessError)

	// i 为了避免被其他包实现
	i()
}

type interceptor struct {
	logger      *zap.Logger
	cache       redis.Repo
	db          mysql.Repo
	userService user.Service
}

func New(logger *zap.Logger, cache redis.Repo, db mysql.Repo) Interceptor {
	return &interceptor{
		logger:      logger,
		cache:       cache,
		db:          db,
		userService: user.New(db, cache),
	}
}

func (i *interceptor) i() {}
