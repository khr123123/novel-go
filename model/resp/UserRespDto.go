package resp

// UserLoginRespDto 用户登录响应
type UserLoginRespDto struct {
	UID      int64  `json:"uid" example:"1"`                     // 用户ID
	NickName string `json:"nickName" example:"小明"`               // 用户昵称
	Token    string `json:"token" example:"xxxxx.yyy.zzz"`       // 用户token
	Avatar   string `json:"avatar" example:"/static/avatar.jpg"` // 用户头像
}

// UserRegisterRespDto 用户注册响应 DTO
type UserRegisterRespDto struct {
	UID      int64  `json:"uid" example:"123" swagger:"desc(用户ID)"`
	Token    string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" swagger:"desc(用户token)"`
	NickName string `json:"nickName" example:"张三" swagger:"desc(用户昵称)"`
	Avatar   string `json:"avatar" example:"https://example.com/avatar.png" swagger:"desc(用户头像)"`
}
