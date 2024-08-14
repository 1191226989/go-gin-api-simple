package captcha

import (
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/redis"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 生成验证码id和图片
	// @Tags API.captcha
	// @Router /api/captcha [get]
	Create() core.HandlerFunc

	// Verify 验证码校验
	// @Tags API.captcha
	// @Router /api/captcha [post]
	Verify() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
	cache  redis.Repo
}

func New(logger *zap.Logger, cache redis.Repo) Handler {
	return &handler{
		logger: logger,
		cache:  cache,
	}
}

func (h *handler) i() {}
