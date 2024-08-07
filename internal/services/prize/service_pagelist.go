package prize

import (
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/mysql"
	"go-gin-api-simple/internal/repository/mysql/prize"
)

type SearchData struct {
	Page     int     // 第几页
	PageSize int     // 每页显示条数
	Name     string  // 奖品名称
	Worth    float64 // 奖品价值
	Content  string  // 奖品描述
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (listData []*prize.Prize, err error) {

	page := searchData.Page
	if page == 0 {
		page = 1
	}

	pageSize := searchData.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

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

	listData, err = qb.
		Limit(pageSize).
		Offset(offset).
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
