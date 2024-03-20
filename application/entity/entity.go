package entity

import "time"

type BaseEntity struct {
	Deleted      bool      `gorm:"column:deleted;default:false;comment:删除标记"`
	Creator      string    `gorm:"column:creator;size:128;comment:创建人"`
	CreatedTime  time.Time `gorm:"column:created_time;comment:创建时间"`
	Modifier     string    `gorm:"column:modifier;size:128;comment:修改人"`
	ModifiedTime time.Time `gorm:"column:modified_time;comment:修改时间"`
}
