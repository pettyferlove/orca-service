package entity

import (
	"time"
)

type BaseEntity struct {
	Id        string    `gorm:"column:id;primaryKey;size:128;comment:数据ID"`
	CreatedBy string    `gorm:"column:created_by;size:128;not null;comment:创建人"`
	CreatedAt time.Time `gorm:"column:created_at;not null;comment:创建时间"`
	UpdatedBy string    `gorm:"column:updated_by;size:128;not null;comment:更新人"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;comment:更新时间"`
}
