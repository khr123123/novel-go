package req

// PageReqDto 分页请求数据格式封装，所有分页请求的结构体都应包含此字段
type PageReqDto struct {
	PageNum  int64 `json:"pageNum"  yaml:"pageNum"`  // 请求页码，默认第 1 页
	PageSize int64 `json:"pageSize" yaml:"pageSize"` // 每页大小，默认每页 10 条
	FetchAll bool  `json:"fetchAll" yaml:"fetchAll"` // 是否查询所有，默认 false；true 时 pageNum 和 pageSize 无效
}

// NewPageReqDto 创建分页请求结构体，设置默认值
func NewPageReqDto() *PageReqDto {
	return &PageReqDto{
		PageNum:  1,
		PageSize: 10,
		FetchAll: false,
	}
}
