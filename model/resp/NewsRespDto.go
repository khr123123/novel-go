package resp

import "time"

// NewsInfoRespDto 新闻信息响应 DTO
type NewsInfoRespDto struct {
	ID           int64     `json:"id" example:"1"`                      // 新闻ID
	CategoryID   int64     `json:"categoryID" example:"101"`            // 类别ID
	CategoryName string    `json:"categoryName" example:"公告"`           // 类别名
	SourceName   string    `json:"sourceName" example:"官方来源"`           // 新闻来源
	Title        string    `json:"title" example:"新功能上线"`               // 新闻标题
	UpdateTime   time.Time `json:"updateTime" example:"2025-06-18"`     // 更新时间（默认 RFC3339 格式）
	Content      string    `json:"content,omitempty" example:"具体内容..."` // 新闻内容
}
