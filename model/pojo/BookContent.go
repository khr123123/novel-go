package pojo

import "time"

// BookContent 小说内容实体
type BookContent struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ChapterID  int64     `gorm:"column:chapter_id" json:"chapterId"`   // 章节ID
	Content    string    `gorm:"type:text" json:"content"`             // 小说章节内容
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
}

// TableName 指定对应表名
func (BookContent) TableName() string {
	return "book_content"
}
