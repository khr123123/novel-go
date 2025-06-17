package pojo

import (
	"time"
)

// UserInfo 用户信息
type UserInfo struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username       string    `gorm:"column:username;type:varchar(255);not null" json:"username"`
	Password       string    `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Salt           string    `gorm:"column:salt;type:varchar(255)" json:"salt"`
	NickName       string    `gorm:"column:nick_name;type:varchar(255)" json:"nickName"`
	UserPhoto      string    `gorm:"column:user_photo;type:varchar(255)" json:"userPhoto"`
	UserSex        int       `gorm:"column:user_sex;type:int" json:"userSex"` // 0-男 1-女
	AccountBalance int64     `gorm:"column:account_balance;type:bigint" json:"accountBalance"`
	Status         int       `gorm:"column:status;type:int" json:"status"` // 用户状态;0-正常
	CreateTime     time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
}

// TableName 指定表名
func (UserInfo) TableName() string {
	return "user_info"
}
