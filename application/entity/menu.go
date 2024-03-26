package entity

import "gorm.io/gorm"

type Menu struct {
	ParentId     string         `gorm:"column:parent_id;size:128;comment:父菜单ID"`
	PermissionId string         `gorm:"column:permission_id;size:128;not null;comment:权限ID"`
	Label        string         `gorm:"column:label;size:255;not null;comment:菜单名称"`
	Icon         string         `gorm:"column:icon;size:255;comment:图标"`
	Path         string         `gorm:"column:path;size:255;not null;comment:路径"`
	Sort         int            `gorm:"column:sort;not null;default:0;comment:排序"`
	Valid        bool           `gorm:"column:valid;not null;default:true;comment:是否有效"`
	Visible      bool           `gorm:"column:visible;not null;default:true;comment:是否可见"`
	DeletedAt    gorm.DeletedAt `gorm:"index;comment:删除时间"`
	BaseEntity
}

func (Menu) TableName() string {
	return "s_menu"
}
