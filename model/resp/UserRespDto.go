package resp

import "time"

// UserLoginRespDto 用户登录响应
type UserLoginRespDto struct {
	UID      int64  `json:"uid,string" example:"1"`              // 用户ID
	NickName string `json:"nickName" example:"小明"`               // 用户昵称
	Token    string `json:"token" example:"xxxxx.yyy.zzz"`       // 用户token
	Avatar   string `json:"avatar" example:"/static/avatar.jpg"` // 用户头像
}

// UserRegisterRespDto 用户注册响应 DTO
type UserRegisterRespDto struct {
	UID      int64  `json:"uid,string" example:"123" swagger:"desc(用户ID)"`
	Token    string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" swagger:"desc(用户token)"`
	NickName string `json:"nickName" example:"张三" swagger:"desc(用户昵称)"`
	Avatar   string `json:"avatar" example:"https://example.com/avatar.png" swagger:"desc(用户头像)"`
}

// UserCommentRespDto 用户评论响应结构体
type UserCommentRespDto struct {
	CommentContent string    `json:"commentContent" example:"这本小说非常精彩！"`                    // 评论内容
	CommentBookPic string    `json:"commentBookPic" example:"https://example.com/book.jpg"` // 评论小说封面
	CommentBook    string    `json:"commentBook" example:"剑破九天"`                            // 评论小说名
	CommentTime    time.Time `json:"commentTime" example:"2023-04-25 12:34:56"`             // 评论时间
}

type UserInfoRespDto struct {
	NickName  string `json:"nickName"`
	UserPhoto string `json:"userPhoto"`
	UserSex   int    `json:"userSex"`
}
