package prize

import "time"

// Prize 奖品表
//
//go:generate gormgen -structs Prize -input .
type Prize struct {
	Id        int32     //
	Name      string    // 奖品名称
	Image     string    // 奖品图片
	Worth     float64   // 奖品价值
	Content   string    // 奖品描述
	IsUsed    int32     // 是否启用 1:是  -1:否
	CreatedAt time.Time `gorm:"time"` // 创建时间
	UpdatedAt time.Time `gorm:"time"` // 更新时间
}
