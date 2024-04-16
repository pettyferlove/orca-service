package constant

// RoleType 角色类型常量
type RoleType string

const (
	// SystemRoleType 系统角色
	SystemRoleType RoleType = "SYSTEM"
	// BusinessRoleType 业务角色
	BusinessRoleType RoleType = "BUSINESS"
)

// RoleTypes 角色类型常量列表
var RoleTypes = []RoleType{SystemRoleType, BusinessRoleType}

// String 转换为字符串
func (r RoleType) String() string {
	return string(r)
}

// Contains 是否包含
func (r RoleType) Contains(role RoleType) bool {
	for _, v := range RoleTypes {
		if v == role {
			return true
		}
	}
	return false
}

// ContainsString 是否包含
func (r RoleType) ContainsString(role string) bool {
	for _, v := range RoleTypes {
		if v.String() == role {
			return true
		}
	}
	return false
}
