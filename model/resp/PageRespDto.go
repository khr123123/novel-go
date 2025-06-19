package resp

// PageRespDto 分页响应数据格式封装，T 是数据类型
type PageRespDto[T any] struct {
	PageNum  int64 `json:"pageNum"`  // 当前页码
	PageSize int64 `json:"pageSize"` // 每页大小
	Total    int64 `json:"total"`    // 总记录数
	List     []T   `json:"list"`     // 分页数据集
}

// NewPageRespDto 创建分页响应对象
func NewPageRespDto[T any](pageNum, pageSize, total int64, list []T) *PageRespDto[T] {
	return &PageRespDto[T]{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    total,
		List:     list,
	}
}

// Pages 计算总页数
func (p *PageRespDto[T]) Pages() int64 {
	if p.PageSize == 0 {
		return 0
	}
	pages := p.Total / p.PageSize
	if p.Total%p.PageSize != 0 {
		pages++
	}
	return pages
}
