package user

import (
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/mysql"
	"go-gin-api-simple/internal/repository/mysql/user"
)

type ModifyData struct {
	Nickname string // 昵称
	Mobile   string // 手机号
}

func (s *service) ModifyPersonalInfo(ctx core.Context, id int32, modifyData *ModifyData) (err error) {
	data := map[string]interface{}{
		"nickname": modifyData.Nickname,
		"mobile":   modifyData.Mobile,
	}

	qb := user.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
