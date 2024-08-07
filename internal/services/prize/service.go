package prize

import (
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/mysql"
	"go-gin-api-simple/internal/repository/mysql/prize"
	"go-gin-api-simple/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, prizeData *CreatePrizeData) (id int32, err error)
	PageList(ctx core.Context, searchData *SearchData) (listData []*prize.Prize, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx core.Context, id int32, used int32) (err error)
	Delete(ctx core.Context, id int32) (err error)
	Detail(ctx core.Context, searchOneData *SearchOneData) (info *prize.Prize, err error)
	Modify(ctx core.Context, id int32, modifyData *ModifyPrizeData) (err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func New(db mysql.Repo, cache redis.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}
