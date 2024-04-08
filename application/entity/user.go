package entity

import (
	"gorm.io/gorm"
	"orca-service/global/security"
	"time"
)

type User struct {
	Username          string              `gorm:"column:username;size:255;not null;comment:用户名"`
	Password          string              `gorm:"column:password;size:255;not null;comment:登录密码"`
	Channel           string              `gorm:"column:channel;size:128;not null;comment:渠道"`
	Status            security.UserStatus `gorm:"column:status;not null;comment:状态"`
	LoginFail         int                 `gorm:"column:login_fail;comment:登录失败次数"`
	LastLoginFailTime time.Time           `gorm:"column:last_login_fail_time;comment:最后登录失败时间"`
	DeletedAt         gorm.DeletedAt      `gorm:"index;comment:删除时间"`
	BaseEntity
	UserInfo UserInfo   `gorm:"foreignKey:user_id"`
	UserRole []UserRole `gorm:"foreignKey:user_id"`
	Roles    []Role     `gorm:"many2many:s_user_role;joinForeignKey:user_id;joinReferences:role_id"`
}

func (User) TableName() string {
	return "s_user"
}
