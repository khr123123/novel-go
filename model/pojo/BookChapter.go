package pojo

import "time"

// BookChapter 小说章节实体
type BookChapter struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	BookID      int64     `gorm:"column:book_id" json:"bookId"`           // 小说ID
	ChapterNum  int       `gorm:"column:chapter_num" json:"chapterNum"`   // 章节号
	ChapterName string    `gorm:"column:chapter_name" json:"chapterName"` // 章节名
	WordCount   int       `gorm:"column:word_count" json:"wordCount"`     // 章节字数
	IsVip       int       `gorm:"column:is_vip" json:"isVip"`             // 是否收费;1-收费 0-免费
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`   // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`   // 更新时间
}

// TableName 指定对应的表名
func (BookChapter) TableName() string {
	return "book_chapter"
}
