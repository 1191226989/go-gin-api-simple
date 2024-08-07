package user

import (
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/pkg/password"
	"go-gin-api-simple/internal/repository/mysql"
	"go-gin-api-simple/internal/repository/mysql/user"
)

func (s *service) ModifyPassword(ctx core.Context, id int32, newPassword string) (err error) {
	data := map[string]interface{}{
		"password": password.GeneratePassword(newPassword),
	}

	qb := user.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	// s.cache.Del(configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id), redis.WithTrace(ctx.Trace()))
	return
}
