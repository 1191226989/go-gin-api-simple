package prize

import (
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/mysql/prize"
)

type CreatePrizeData struct {
	Name    string  // 奖品名称
	Image   string  // 奖品图片
	Worth   float64 // 奖品价值
	IsUsed  int32   // 是否启用
	Content string  // 奖品描述
}

func (s *service) Create(ctx core.Context, prizeData *CreatePrizeData) (id int32, err error) {
	model := prize.NewModel()
	model.Name = prizeData.Name
	model.Image = prizeData.Image
	model.Worth = prizeData.Worth
	model.Content = prizeData.Content
	model.IsUsed = 1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
