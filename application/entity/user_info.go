package entity

import "time"

type UserInfo struct {
	Id       string    `gorm:"column:id;primaryKey;size:128;comment:数据ID"`
	UserId   string    `gorm:"column:user_id;size:128;comment:用户ID"`
	Nickname string    `gorm:"column:nickname;size:128;comment:用户昵称"`
	FullName string    `gorm:"column:full_name;size:128;comment:姓名"`
	Phone    string    `gorm:"column:phone;size:128;comment:手机号码"`
	Gender   int8      `gorm:"column:gender;comment:性别"`
	Birthday time.Time `gorm:"column:birthday;comment:生日"`
	Avatar   string    `gorm:"column:avatar;size:1000;comment:用户头像"`
	Email    string    `gorm:"column:email;size:128;comment:电子邮件"`
	Address  string    `gorm:"column:address;size:400;comment:居住地址"`
	BaseEntity
}

func (UserInfo) TableName() string {
	return "s_user_info"
}
