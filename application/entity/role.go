package entity

import (
	"gorm.io/gorm"
	"orca-service/application/constant"
)

type Role struct {
	Role         string            `gorm:"column:role;size:255;not null;comment:角色"`
	RoleName     string            `gorm:"column:role_name;size:255;not null;comment:角色名称"`
	RoleType     constant.RoleType `gorm:"column:role_type;size:255;not null;comment:角色类型"`
	RoleCategory string            `gorm:"column:role_category;size:255;comment:角色分类"`
	Remark       string            `gorm:"column:remark;size:2048;comment:备注"`
	Valid        bool              `gorm:"column:valid;not null;default:true;comment:是否有效"`
	DeletedAt    gorm.DeletedAt    `gorm:"index;comment:删除时间"`
	BaseEntity
	RolePermission []RolePermission `gorm:"foreignKey:role_id"`
	RoleMenu       []RoleMenu       `gorm:"foreignKey:role_id"`
	UserRole       []UserRole       `gorm:"foreignKey:role_id"`
	Users          []User           `gorm:"many2many:s_user_role;joinForeignKey:role_id;joinReferences:user_id"`
	Permissions    []Permission     `gorm:"many2many:s_role_permission;joinForeignKey:role_id;joinReferences:permission_id"`
}

func (Role) TableName() string {
	return "s_role"
}
