package entity

import "time"

type UserInfo struct {
	Id       string    `json:"id" gorm:"primaryKey;size:128;comment:数据ID"`
	UserId   string    `json:"user_id" gorm:"size:128;comment:用户ID"`
	UserName string    `json:"user_name" gorm:"size:128;comment:用户名"`
	NickName string    `json:"nick_name" gorm:"size:128;comment:用户昵称"`
	Gender   int8      `json:"gender" gorm:"comment:性别"`
	Birthday time.Time `json:"birthday" gorm:"comment:生日"`
	Avatar   string    `json:"avatar" gorm:"size:1000;comment:用户头像"`
	Email    string    `json:"email" gorm:"size:128;comment:电子邮件"`
	Address  string    `json:"address" gorm:"size:400;comment:居住地址"`
	BaseEntity
}

func (UserInfo) TableName() string {
	return "s_user_info"
}
