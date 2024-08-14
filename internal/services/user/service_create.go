package user

import (
	"time"

	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/mysql/user"
)

type CreateUserData struct {
	Username  string    // 用户名
	Password  string    // 密码
	Nickname  string    // 昵称
	Mobile    string    // 手机号
	IsUsed    int32     // 是否启用 1:是  -1:否
	IsDeleted int32     // 是否删除 1:是  -1:否
	CreatedAt time.Time `gorm:"time"` // 创建时间
	UpdatedAt time.Time `gorm:"time"` // 更新时间
}

func (s *service) Create(ctx core.Context, createData *CreateUserData) (id int32, err error) {
	model := user.NewModel()
	model.Username = createData.Username
	model.Password = createData.Password
	model.Nickname = createData.Nickname
	model.Mobile = createData.Mobile
	model.IsUsed = createData.IsUsed
	model.IsDeleted = createData.IsDeleted

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}

	return
}
