package entity

// SecurityRule 安全规则（表结构不允许调整）
type SecurityRule struct {
	ID    uint   `gorm:"primaryKey;column:id;type:int(11) unsigned;not null;autoIncrement;comment:主键ID"`
	Ptype string `gorm:"size:512;column:ptype;type:varchar(512);not null;comment:策略类型"`
	V0    string `gorm:"size:512;column:v0;type:varchar(512);not null;comment:租户"`
	V1    string `gorm:"size:512;column:v1;type:varchar(512);not null;comment:主体"`
	V2    string `gorm:"size:512;column:v2;type:varchar(512);not null;comment:客体"`
	V3    string `gorm:"size:512;column:v3;type:varchar(512);not null;comment:动作"`
	V4    string `gorm:"size:512;column:v4;type:varchar(512);not null;comment:服务"`
	V5    string `gorm:"size:512;column:v5;type:varchar(512);not null;comment:效果"`
}

func (SecurityRule) TableName() string {
	return "s_security_rule"
}
