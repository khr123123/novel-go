package resp

// HomeBookRespDto 表示首页小说推荐的响应结构体
type HomeBookRespDto struct {
	Type       int    `json:"type"`       // 类型：0-轮播图，1-顶部栏，2-本周强推，3-热门推荐，4-精品推荐
	BookID     int64  `json:"bookId"`     // 小说ID
	PicUrl     string `json:"picUrl"`     // 小说封面地址
	BookName   string `json:"bookName"`   // 小说名
	AuthorName string `json:"authorName"` // 作家名
	BookDesc   string `json:"bookDesc"`   // 书籍描述
}

// HomeFriendLinkRespDto 表示首页友情链接的响应结构体
type HomeFriendLinkRespDto struct {
	LinkName string `json:"linkName"` // 链接名
	LinkUrl  string `json:"linkUrl"`  // 链接url
}
