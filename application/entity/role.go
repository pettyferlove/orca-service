package entity

import "gorm.io/gorm"

type Role struct {
	Role         string         `gorm:"column:role;size:255;not null;comment:角色"`
	RoleName     string         `gorm:"column:role_name;size:255;not null;comment:角色名称"`
	RoleType     string         `gorm:"column:role_type;size:255;not null;comment:角色类型"`
	RoleCategory string         `gorm:"column:role_category;size:255;comment:角色分类"`
	Remark       string         `gorm:"column:remark;size:2048;comment:备注"`
	Valid        bool           `gorm:"column:valid;not null;default:true;comment:是否有效"`
	DeletedAt    gorm.DeletedAt `gorm:"index;comment:删除时间"`
	BaseEntity
}

func (Role) TableName() string {
	return "s_role"
}
