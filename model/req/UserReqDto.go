package req

type UserLoginReqDto struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserRegisterReqDto struct {
	Username  string `json:"username" binding:"required,len=11,numeric"` // 手机号，必须11位数字
	Password  string `json:"password" binding:"required"`                // 密码，必填
	VelCode   string `json:"velCode" binding:"required,len=5,numeric"`   // 验证码，5位数字
	SessionId string `json:"sessionId" binding:"required,len=20"`        // sessionId，20位字符串
}
type UserCommentReqDto struct {
	UserId         int64  `json:"userId,omitempty"`                                 // 可选用户ID，指针表示可为空
	BookId         int64  `json:"bookId" binding:"required"`                        // 小说ID，必填
	CommentContent string `json:"commentContent" binding:"required,min=10,max=512"` // 评论内容，必填，长度10~512
}
