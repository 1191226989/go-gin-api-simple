package prize

import (
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/mysql"
	"go-gin-api-simple/internal/repository/mysql/prize"
)

type ModifyPrizeData struct {
	Name    string  // 奖品名称
	Image   string  // 奖品图片
	Worth   float64 // 奖品价值
	Content string  // 奖品描述
	IsUsed  int32   // 是否启用
}

func (s *service) Modify(ctx core.Context, id int32, modifyData *ModifyPrizeData) (err error) {
	data := map[string]interface{}{
		"name":    modifyData.Name,
		"image":   modifyData.Image,
		"worth":   modifyData.Worth,
		"content": modifyData.Content,
		"is_used": modifyData.IsUsed,
	}

	qb := prize.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
