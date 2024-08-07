package user

import (
	"go-gin-api-simple/configs"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/mysql"
	"go-gin-api-simple/internal/repository/redis"
	"go-gin-api-simple/internal/services/user"
	"go-gin-api-simple/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Signup 用户注册
	// @Tags API.user
	// @Router /api/signup [post]
	Signup() core.HandlerFunc

	// Login 用户登录
	// @Tags API.user
	// @Router /api/login [post]
	Login() core.HandlerFunc

	// Logout 用户登出
	// @Tags API.user
	// @Router /api/user/logout [post]
	Logout() core.HandlerFunc

	// ModifyPassword 修改密码
	// @Tags API.user
	// @Router /api/user/modify_password [patch]
	ModifyPassword() core.HandlerFunc

	// Detail 个人信息
	// @Tags API.user
	// @Router /api/user/info [get]
	Detail() core.HandlerFunc

	// ModifyPersonalInfo 修改个人信息
	// @Tags API.user
	// @Router /api/user/modify_personal_info [patch]
	ModifyPersonalInfo() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       redis.Repo
	hashids     hash.Hash
	userService user.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		userService: user.New(db, cache),
	}
}

func (h *handler) i() {}
