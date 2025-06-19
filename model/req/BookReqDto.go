package req

// BookAddReqDto 小说发布请求 DTO
type BookAddReqDto struct {
	WorkDirection int64  `json:"workDirection" binding:"required,oneof=0 1" example:"0"`             // 作品方向;0-男频 1-女频
	CategoryId    int64  `json:"categoryId" binding:"required" example:"1001"`                       // 类别ID
	CategoryName  string `json:"categoryName" binding:"required" example:"玄幻"`                       // 类别名
	PicUrl        string `json:"picUrl" binding:"required,url" example:"http://example.com/pic.jpg"` // 小说封面地址，简单验证url格式
	BookName      string `json:"bookName" binding:"required" example:"我的小说"`                         // 小说名
	BookDesc      string `json:"bookDesc" binding:"required" example:"小说简介内容"`                       // 书籍描述
	IsVip         int    `json:"isVip" binding:"required,oneof=0 1" example:"0"`                     // 是否收费;1-收费 0-免费
}
