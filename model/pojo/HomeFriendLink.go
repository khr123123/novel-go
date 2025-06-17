package pojo

import "time"

// HomeFriendLink 对应 home_friend_link 表
type HomeFriendLink struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	LinkName   string    `gorm:"column:link_name"`
	LinkUrl    string    `gorm:"column:link_url"`
	Sort       int       `gorm:"column:sort"`
	IsOpen     int       `gorm:"column:is_open"` // 0-关闭 1-开启
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

func (HomeFriendLink) TableName() string {
	return "home_friend_link"
}
