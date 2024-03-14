package model

import "time"

type User struct {
	Id                string    `json:"id" gorm:"primaryKey;size:128;comment:数据ID"`
	LoginName         string    `json:"login_name" gorm:"size:255;comment:登录名"`
	Password          string    `json:"password" gorm:"size:255;comment:登录密码"`
	TenantId          string    `json:"tenant_id" gorm:"size:128;comment:租户ID"`
	Channel           string    `json:"channel" gorm:"size:128;comment:渠道"`
	Status            int8      `json:"status" gorm:"comment:状态"`
	LoginFail         int       `json:"login_fail" gorm:"comment:登录失败次数"`
	LastLoginFailTime time.Time `json:"last_login_fail_time" gorm:"comment:最后登录失败时间"`
	BaseModel
}

func (User) TableName() string {
	return "s_user"
}
