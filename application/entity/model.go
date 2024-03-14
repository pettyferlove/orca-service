package entity

import "time"

type BaseEntity struct {
	DelFlag    bool      `json:"del_flag" gorm:"comment:删除标记"`
	Creator    string    `json:"creator" gorm:"size:128;comment:创建人"`
	CreateTime time.Time `json:"create_time" gorm:"comment:创建时间"`
	Modifier   string    `json:"modifier" gorm:"size:128;comment:修改人"`
	ModifyTime time.Time `json:"modify_time" gorm:"comment:修改时间"`
}
