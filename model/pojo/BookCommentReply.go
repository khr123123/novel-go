package pojo

import "time"

// BookCommentReply 小说评论回复实体
type BookCommentReply struct {
	ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CommentID    int64     `gorm:"column:comment_id" json:"commentId"`       // 评论ID
	UserID       int64     `gorm:"column:user_id" json:"userId"`             // 回复用户ID
	ReplyContent string    `gorm:"column:reply_content" json:"replyContent"` // 回复内容
	AuditStatus  int       `gorm:"column:audit_status" json:"auditStatus"`   // 审核状态;0-待审核 1-审核通过 2-审核不通过
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`     // 更新时间
}

// TableName 显式指定表名
func (BookCommentReply) TableName() string {
	return "book_comment_reply"
}
