package pojo

import "time"

// BookComment 小说评论实体
type BookComment struct {
	ID             int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	BookID         int64     `gorm:"column:book_id" json:"bookId"`                 // 评论小说ID
	UserID         int64     `gorm:"column:user_id" json:"userId"`                 // 评论用户ID
	CommentContent string    `gorm:"column:comment_content" json:"commentContent"` // 评价内容
	ReplyCount     int       `gorm:"column:reply_count" json:"replyCount"`         // 回复数量
	AuditStatus    int       `gorm:"column:audit_status" json:"auditStatus"`       // 审核状态;0-待审核 1-审核通过 2-审核不通过
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`         // 创建时间
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`         // 更新时间
}

// TableName 显式指定表名
func (BookComment) TableName() string {
	return "book_comment"
}
