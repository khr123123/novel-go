package req

type UserRegisterReqDto struct {
	Username  string `json:"username" binding:"required,len=11,numeric"` // 手机号，必须11位数字
	Password  string `json:"password" binding:"required"`                // 密码，必填
	VelCode   string `json:"velCode" binding:"required,len=5,numeric"`   // 验证码，5位数字
	SessionId string `json:"sessionId" binding:"required,len=20"`        // sessionId，20位字符串
}
