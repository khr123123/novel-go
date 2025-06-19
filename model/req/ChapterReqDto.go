package req

// ChapterUpdateReqDto 章节发布 请求结构体
type ChapterUpdateReqDto struct {
	ChapterName    string `json:"chapterName" binding:"required" example:"第一章 开始"`                       // 章节名，必填
	ChapterContent string `json:"chapterContent" binding:"required,min=50" example:"这是章节的内容，至少50个字符..."` // 章节内容，必填，最小长度50
	IsVip          int    `json:"isVip" binding:"required,oneof=0 1" example:"0"`                        // 是否收费; 1-收费 0-免费，必填
}

// ChapterAddReqDto 章节发布请求 DTO
type ChapterAddReqDto struct {
	BookId         int64  `json:"bookId" binding:"required" example:"123"`                         // 小说ID，必填
	ChapterName    string `json:"chapterName" binding:"required" example:"第一章 起始篇"`                // 章节名，必填，非空
	ChapterContent string `json:"chapterContent" binding:"required,min=50" example:"这里是章节正文内容..."` // 章节内容，必填，最小长度50
	IsVip          int    `json:"isVip" binding:"required,oneof=0 1" example:"0"`                  // 是否收费，0-免费，1-收费
}
