package pojo

import "time"

// NewsInfo 对应 news_info 表
type NewsInfo struct {
	ID           int64     `gorm:"primaryKey;column:id" json:"id"`                       // 主键
	CategoryID   int64     `gorm:"column:category_id" json:"category_id"`                // 类别ID
	CategoryName string    `gorm:"column:category_name" json:"category_name"`            // 类别名
	SourceName   string    `gorm:"column:source_name" json:"source_name"`                // 新闻来源
	Title        string    `gorm:"column:title" json:"title"`                            // 新闻标题
	CreateTime   time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"` // 更新时间
}

// TableName 显式声明表名
func (NewsInfo) TableName() string {
	return "news_info"
}
