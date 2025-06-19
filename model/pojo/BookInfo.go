package pojo

import (
	"time"
)

// BookInfo 小说信息表结构体
type BookInfo struct {
	ID                    int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	WorkDirection         int       `gorm:"column:work_direction" json:"workDirection"`                   // 作品方向;0-男频 1-女频
	CategoryID            int64     `gorm:"column:category_id" json:"categoryId"`                         // 类别ID
	CategoryName          string    `gorm:"column:category_name" json:"categoryName"`                     // 类别名
	PicURL                string    `gorm:"column:pic_url" json:"picUrl"`                                 // 小说封面地址
	BookName              string    `gorm:"column:book_name" json:"bookName"`                             // 小说名
	AuthorID              int64     `gorm:"column:author_id" json:"authorId"`                             // 作家ID
	AuthorName            string    `gorm:"column:author_name" json:"authorName"`                         // 作家名
	BookDesc              string    `gorm:"column:book_desc" json:"bookDesc"`                             // 书籍描述
	Score                 int       `gorm:"column:score" json:"score"`                                    // 评分; 总分10，真实评分 = score/10
	BookStatus            int       `gorm:"column:book_status" json:"bookStatus"`                         // 书籍状态;0-连载中 1-已完结
	VisitCount            int64     `gorm:"column:visit_count" json:"visitCount"`                         // 点击量
	WordCount             int       `gorm:"column:word_count" json:"wordCount"`                           // 总字数
	CommentCount          int       `gorm:"column:comment_count" json:"commentCount"`                     // 评论数
	LastChapterID         int64     `gorm:"column:last_chapter_id" json:"lastChapterId"`                  // 最新章节ID
	LastChapterName       string    `gorm:"column:last_chapter_name" json:"lastChapterName"`              // 最新章节名
	LastChapterUpdateTime time.Time `gorm:"column:last_chapter_update_time" json:"lastChapterUpdateTime"` // 最新章节更新时间
	IsVip                 int       `gorm:"column:is_vip" json:"isVip"`                                   // 是否收费;1-收费 0-免费
	CreateTime            time.Time `gorm:"column:create_time" json:"createTime"`                         // 创建时间
	UpdateTime            time.Time `gorm:"column:update_time" json:"updateTime"`                         // 更新时间
}

// TableName 自定义表名，默认会用结构体名的小写加复数形式
func (BookInfo) TableName() string {
	return "book_info"
}
