package entity

import (
	"gorm.io/gorm"
	"orca-service/application/constant"
)

type Permission struct {
	ParentId           string                    `gorm:"column:parent_id;size:128;comment:父权限ID"`
	Permission         string                    `gorm:"column:permission;size:512;not null;comment:权限"`
	PermissionMetadata string                    `gorm:"column:permission_metadata;size:512;comment:权限元数据"`
	PermissionMethod   constant.PermissionMethod `gorm:"column:permission_method;size:128;comment:权限方法"`
	PermissionName     string                    `gorm:"column:permission_name;size:512;not null;comment:权限名称"`
	PermissionType     constant.PermissionType   `gorm:"column:permission_type;size:128;not null;comment:权限类型"`
	PermissionUrl      string                    `gorm:"column:permission_url;size:512;comment:权限URL"`
	Remark             string                    `gorm:"column:remark;size:1024;comment:备注"`
	Sort               int                       `gorm:"column:sort;not null;default:0;comment:排序"`
	Valid              bool                      `gorm:"column:valid;not null;default:true;comment:是否有效"`
	DeletedAt          gorm.DeletedAt            `gorm:"index;comment:删除时间"`
	BaseEntity
	Menu  *Menu  `gorm:"foreignKey:permission_id"`
	Roles []Role `gorm:"many2many:s_role_permission;joinForeignKey:permission_id;joinReferences:role_id"`
}

func (Permission) TableName() string {
	return "s_permission"
}
