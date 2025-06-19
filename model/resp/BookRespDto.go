package resp

import "time"

// BookRankRespDto 小说排行榜 响应结构体
type BookRankRespDto struct {
	ID                    int64     `json:"id,string"`             // 小说ID
	CategoryID            int64     `json:"categoryId,string"`     // 类别ID
	CategoryName          string    `json:"categoryName"`          // 类别名
	PicURL                string    `json:"picUrl"`                // 小说封面地址
	BookName              string    `json:"bookName"`              // 小说名
	AuthorName            string    `json:"authorName"`            // 作家名
	BookDesc              string    `json:"bookDesc"`              // 书籍描述
	WordCount             int       `json:"wordCount"`             // 总字数
	LastChapterName       string    `json:"lastChapterName"`       // 最新章节名
	LastChapterUpdateTime time.Time `json:"lastChapterUpdateTime"` // 最新章节更新时间
}

// BookInfoRespDto 小说信息 响应结构体
type BookInfoRespDto struct {
	ID              int64     `json:"id,string"`             // 小说ID
	CategoryID      int64     `json:"categoryId,string"`     // 类别ID
	CategoryName    string    `json:"categoryName"`          // 类别名
	PicURL          string    `json:"picUrl"`                // 小说封面地址
	BookName        string    `json:"bookName"`              // 小说名
	AuthorID        int64     `json:"authorId,string"`       // 作家ID
	AuthorName      string    `json:"authorName"`            // 作家名
	BookDesc        string    `json:"bookDesc"`              // 书籍描述
	BookStatus      int       `json:"bookStatus"`            // 书籍状态;0-连载中 1-已完结
	VisitCount      int64     `json:"visitCount"`            // 点击量
	WordCount       int       `json:"wordCount"`             // 总字数
	CommentCount    int       `json:"commentCount"`          // 评论数
	FirstChapterId  int64     `json:"firstChapterId,string"` // 首章节ID
	LastChapterId   int64     `json:"lastChapterId,string"`  // 最新章节ID
	LastChapterName string    `json:"lastChapterName"`       // 最新章节名
	UpdateTime      time.Time `json:"updateTime"`            // 更新时间
}

// BookContentAboutRespDto 小说内容相关 响应结构体
type BookContentAboutRespDto struct {
	BookInfo    BookInfoRespDto    `json:"bookInfo"`    // 小说信息
	ChapterInfo BookChapterRespDto `json:"chapterInfo"` // 章节信息
	BookContent string             `json:"bookContent"` // 章节内容
}

// BookCommentRespDto 小说评论响应 DTO
type BookCommentRespDto struct {
	CommentTotal int64         `json:"commentTotal"` // 评论总数
	Comments     []CommentInfo `json:"comments"`     // 评论列表
}

// CommentInfo 评论信息
type CommentInfo struct {
	ID               int64     `json:"id,string"`            // 评论ID
	CommentContent   string    `json:"commentContent"`       // 评论内容
	CommentUser      string    `json:"commentUser"`          // 评论用户
	CommentUserId    int64     `json:"commentUserId,string"` // 评论用户ID
	CommentUserPhoto string    `json:"commentUserPhoto"`     // 评论用户头像
	CommentTime      time.Time `json:"commentTime"`          // 评论时间
}

// BookCategoryRespDto 小说类别 响应结构体
type BookCategoryRespDto struct {
	ID   int64  `json:"id,string"` // 类别ID
	Name string `json:"name"`      // 类别名
}
