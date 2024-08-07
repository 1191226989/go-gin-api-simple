package prize

import (
	"go-gin-api-simple/configs"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/mysql"
	"go-gin-api-simple/internal/repository/redis"
	"go-gin-api-simple/internal/services/prize"
	"go-gin-api-simple/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// List 奖品列表
	// @Tags API.prize
	// @Router /api/prize [get]
	List() core.HandlerFunc

	// Detail 奖品详情
	// @Tags API.prize
	// @Router /api/prize/detail [get]
	Detail() core.HandlerFunc

	// Modify 编辑奖品
	// @Tags API.prize
	// @Router /api/prize/{id} [post]
	Modify() core.HandlerFunc

	// Create 新增奖品
	// @Tags API.prize
	// @Router /api/prize [post]
	Create() core.HandlerFunc

	// Delete 删除奖品
	// @Tags API.prize
	// @Router /api/prize/{id} [delete]
	Delete() core.HandlerFunc

	// UpdateUsed 更新奖品为启用/禁用
	// @Tags API.prize
	// @Router /api/prize/used [patch]
	UpdateUsed() core.HandlerFunc
}

type handler struct {
	logger       *zap.Logger
	cache        redis.Repo
	hashids      hash.Hash
	prizeService prize.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:       logger,
		cache:        cache,
		hashids:      hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		prizeService: prize.New(db, cache),
	}
}

func (h *handler) i() {}
