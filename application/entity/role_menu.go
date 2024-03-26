package entity

import "time"

type RoleMenu struct {
	RoleId     string    `gorm:"column:role_id;size:128;not null;comment:角色ID"`
	MenuId     string    `gorm:"column:menu_id;size:128;not null;comment:菜单ID"`
	ValidBegin time.Time `gorm:"column:valid_begin;comment:有效开始时间"`
	ValidEnd   time.Time `gorm:"column:valid_end;comment:有效结束时间"`
	BaseEntity
}

func (RoleMenu) TableName() string {
	return "s_role_menu"
}
