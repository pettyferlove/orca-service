package entity

import "time"

type UserRole struct {
	UserId     string    `gorm:"column:user_id;size:128;not null;comment:用户ID"`
	RoleId     string    `gorm:"column:role_id;size:128;not null;comment:角色ID"`
	ValidBegin time.Time `gorm:"column:valid_begin;comment:有效开始时间"`
	ValidEnd   time.Time `gorm:"column:valid_end;comment:有效结束时间"`
	BaseEntity
	User User `gorm:"foreignKey:user_id"`
	Role Role `gorm:"foreignKey:role_id"`
}

func (UserRole) TableName() string {
	return "s_user_role"
}
