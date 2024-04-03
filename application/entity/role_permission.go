package entity

import "time"

type RolePermission struct {
	RoleId       string    `gorm:"column:role_id;size:128;not null;comment:角色ID"`
	PermissionId string    `gorm:"column:permission_id;size:128;not null;comment:权限ID"`
	ValidBegin   time.Time `gorm:"column:valid_begin;comment:有效开始时间"`
	ValidEnd     time.Time `gorm:"column:valid_end;comment:有效结束时间"`
	BaseEntity
	Role       Role       `gorm:"foreignKey:role_id"`
	Permission Permission `gorm:"foreignKey:permission_id"`
}

func (RolePermission) TableName() string {
	return "s_role_permission"
}
