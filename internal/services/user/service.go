package user

import (
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/mysql"
	"go-gin-api-simple/internal/repository/mysql/user"
	"go-gin-api-simple/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, createData *CreateUserData) (id int32, err error)
	Detail(ctx core.Context, searchOneData *SearchOneData) (info *user.User, err error)
	ModifyPassword(ctx core.Context, id int32, newPassword string) (err error)
	ModifyPersonalInfo(ctx core.Context, id int32, modifyData *ModifyData) (err error)
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
