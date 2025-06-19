package pojo

import "time"

// NewsContent 对应 news_content 表
type NewsContent struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"`
	NewsID     int64     `gorm:"column:news_id" json:"news_id"`
	Content    string    `gorm:"column:content" json:"content"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

// TableName 显式声明表名
func (NewsContent) TableName() string {
	return "news_content"
}
