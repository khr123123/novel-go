package resp

// UserRegisterRespDto 用户注册响应 DTO
type UserRegisterRespDto struct {
	UID      int64  `json:"uid" example:"123" swagger:"desc(用户ID)"`
	Token    string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" swagger:"desc(用户token)"`
	NickName string `json:"nickName" example:"张三" swagger:"desc(用户昵称)"`
	Avatar   string `json:"avatar" example:"https://example.com/avatar.png" swagger:"desc(用户头像)"`
}
