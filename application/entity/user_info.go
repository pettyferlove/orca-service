package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	UserId    string         `gorm:"column:user_id;size:128;not null;comment:用户ID"`
	Name      string         `gorm:"column:name;size:128;not null;comment:姓名"`
	Nickname  string         `gorm:"column:nickname;size:128;comment:用户昵称"`
	Phone     string         `gorm:"column:phone;size:128;comment:手机号码"`
	Gender    int8           `gorm:"column:gender;comment:性别"`
	Birthday  time.Time      `gorm:"column:birthday;comment:生日"`
	Avatar    string         `gorm:"column:avatar;size:1000;comment:用户头像"`
	Email     string         `gorm:"column:email;size:128;comment:电子邮件"`
	Address   string         `gorm:"column:address;size:400;comment:居住地址"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间"`
	BaseEntity
}

func (UserInfo) TableName() string {
	return "s_user_info"
}
