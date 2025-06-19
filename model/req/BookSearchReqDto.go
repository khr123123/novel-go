package req

// BookSearchReqDto 小说搜索请求 DTO
type BookSearchReqDto struct {
	Keyword       string  `form:"keyword"`       // 搜索关键字
	WorkDirection *int    `form:"workDirection"` // 作品方向 可选
	Category      *string `form:"category"`      // 分类ID 可选
	IsVip         *int    `form:"isVip"`         // 是否收费 0/1 可选
	BookStatus    *int    `form:"bookStatus"`    // 小说状态 0/1 可选
	WordCountMin  *int    `form:"wordCountMin"`  // 字数最小 可选
	WordCountMax  *int    `form:"wordCountMax"`  // 字数最大 可选
	UpdateTime    *string `form:"updateTime"`    // 最小更新时间 可选 (时间字符串, 也可以改成 time.Time 指针根据需要)
	Sort          string  `form:"sort"`          // 排序字段
	PageNum       int64   `form:"pageNum,default=1"`
	PageSize      int64   `form:"pageSize,default=10"`
}
