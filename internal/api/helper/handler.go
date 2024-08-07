package helper

import (
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/mysql"
	"go-gin-api-simple/internal/repository/redis"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Md5 加密
	// @Tags Helper
	// @Router /helper/md5/{str} [get]
	Md5() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
	db     mysql.Repo
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger: logger,
		db:     db,
	}
}

func (h *handler) i() {}
