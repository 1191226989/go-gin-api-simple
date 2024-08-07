package prize

import (
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/mysql"
	"go-gin-api-simple/internal/repository/mysql/prize"
)

func (s *service) Delete(ctx core.Context, id int32) (err error) {
	qb := prize.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Delete(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return err
	}

	return
}
