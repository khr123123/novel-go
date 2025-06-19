package resp

import "time"

// BookChapterAboutRespDto 小说章节相关 响应结构体
type BookChapterAboutRespDto struct {
	ChapterInfo    BookChapterRespDto `json:"chapterInfo"`    // 章节信息
	ChapterTotal   int64              `json:"chapterTotal"`   // 章节总数
	ContentSummary string             `json:"contentSummary"` // 内容概要（30字）
}

// BookChapterRespDto 小说章节 响应结构体
type BookChapterRespDto struct {
	ID                int64     `json:"id,string" example:"123"`                         // 章节ID
	BookID            int64     `json:"bookId,string" example:"456"`                     // 小说ID
	ChapterNum        int       `json:"chapterNum" example:"1"`                          // 章节号
	ChapterName       string    `json:"chapterName" example:"第一章 开始"`                    // 章节名
	ChapterWordCount  int       `json:"chapterWordCount" example:"3500"`                 // 章节字数
	ChapterUpdateTime time.Time `json:"chapterUpdateTime" example:"2023-04-01 18:00:00"` // 更新时间
	IsVip             int       `json:"isVip" example:"0"`                               // 是否收费;1-收费 0-免费
}

// ChapterContentRespDto 小说内容响应 DTO
type ChapterContentRespDto struct {
	ChapterName    string `json:"chapterName"`    // 章节名
	ChapterContent string `json:"chapterContent"` // 章节内容
	IsVip          int    `json:"isVip"`          // 是否收费;1-收费 0-免费
}
