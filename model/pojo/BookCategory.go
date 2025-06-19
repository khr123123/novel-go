package pojo

import "time"

// BookCategory 小说类别实体
type BookCategory struct {
	ID            int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	WorkDirection int       `gorm:"column:work_direction" json:"workDirection"` // 作品方向;0-男频 1-女频
	Name          string    `gorm:"column:name" json:"name"`                    // 类别名
	Sort          int       `gorm:"column:sort" json:"sort"`                    // 排序
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`       // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`       // 更新时间
}

func (BookCategory) TableName() string {
	return "book_category"
}
