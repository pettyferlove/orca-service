package entity

import "time"

type BaseEntity struct {
	Deleted    bool      `json:"deleted" gorm:"deleted;default:false;comment:删除标记"`
	Creator    string    `json:"creator" gorm:"creator;size:128;comment:创建人"`
	CreateTime time.Time `json:"created_time" gorm:"created_time;comment:创建时间"`
	Modifier   string    `json:"modifier" gorm:"modifier;size:128;comment:修改人"`
	ModifyTime time.Time `json:"modified_time" gorm:"modified_time;comment:修改时间"`
}
