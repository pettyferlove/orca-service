package entity

import (
	"time"
)

type BaseEntity struct {
	Id        string `gorm:"column:id;primaryKey;size:128;comment:数据ID"`
	CreatedBy string `gorm:"column:created_by;size:128;comment:创建人"`
	CreatedAt time.Time
	UpdatedBy string `gorm:"column:updated_by;size:128;comment:更新人"`
	UpdatedAt time.Time
}
