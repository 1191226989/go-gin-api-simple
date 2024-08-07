package prize

import (
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/mysql"
	"go-gin-api-simple/internal/repository/mysql/prize"
)

func (s *service) PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error) {
	qb := prize.NewQueryBuilder()

	if searchData.Name != "" {
		qb.WhereName(mysql.EqualPredicate, searchData.Name)
	}

	if searchData.Worth != 0 {
		qb.WhereWorth(mysql.EqualPredicate, searchData.Worth)
	}

	if searchData.Content != "" {
		qb.WhereContent(mysql.EqualPredicate, searchData.Content)
	}

	total, err = qb.Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}

	return
}
