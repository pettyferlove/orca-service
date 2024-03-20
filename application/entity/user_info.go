package entity

import "time"

type UserInfo struct {
	Id       string    `gorm:"id;primaryKey;size:128;comment:数据ID"`
	UserId   string    `gorm:"user_id;size:128;comment:用户ID"`
	Username string    `gorm:"username;size:128;comment:用户名"`
	Nickname string    `gorm:"nickname;size:128;comment:用户昵称"`
	Gender   int8      `gorm:"gender;comment:性别"`
	Birthday time.Time `gorm:"birthday;comment:生日"`
	Avatar   string    `gorm:"avatar;size:1000;comment:用户头像"`
	Email    string    `gorm:"email;size:128;comment:电子邮件"`
	Address  string    `gorm:"address;size:400;comment:居住地址"`
	BaseEntity
}

func (UserInfo) TableName() string {
	return "s_user_info"
}
