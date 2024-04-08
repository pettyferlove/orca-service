package role

// Type 角色类型常量
type Type string

const (
	// SYSTEM 系统
	SYSTEM Type = "SYSTEM"
	// BUSINESS 业务
	BUSINESS Type = "BUSINESS"
)

// Types 角色类型常量列表
var Types = []Type{SYSTEM, BUSINESS}

// String 转换为字符串
func (r Type) String() string {
	return string(r)
}

// Contains 是否包含
func (r Type) Contains(role Type) bool {
	for _, v := range Types {
		if v == role {
			return true
		}
	}
	return false
}

// ContainsString 是否包含
func (r Type) ContainsString(role string) bool {
	for _, v := range Types {
		if v.String() == role {
			return true
		}
	}
	return false
}
