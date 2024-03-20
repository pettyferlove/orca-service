package entity

import "time"

type User struct {
	Id                string    `gorm:"id;primaryKey;size:128;comment:数据ID"`
	LoginName         string    `gorm:"login_name;size:255;comment:登录名"`
	Password          string    `gorm:"password;size:255;comment:登录密码"`
	TenantId          string    `gorm:"tenant_id;size:128;comment:租户ID"`
	Channel           string    `gorm:"channel;size:128;comment:渠道"`
	Status            int8      `gorm:"status;comment:状态"`
	LoginFail         int       `gorm:"login_fail;comment:登录失败次数"`
	LastLoginFailTime time.Time `gorm:"last_login_fail_time;comment:最后登录失败时间"`
	BaseEntity
}

func (User) TableName() string {
	return "s_user"
}
