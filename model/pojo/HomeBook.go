package pojo

import "time"

// HomeBook 对应 home_book 表
type HomeBook struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	Type       int       `gorm:"column:type"`    // 推荐类型
	Sort       int       `gorm:"column:sort"`    // 排序
	BookID     int64     `gorm:"column:book_id"` // 推荐小说ID
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

func (HomeBook) TableName() string {
	return "home_book"
}
