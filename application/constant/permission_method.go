package constant

// PermissionMethod 权限方法常量
type PermissionMethod string

const (
	// GetPermissionMethod 查询
	GetPermissionMethod PermissionMethod = "GET"
	// PostPermissionMethod 创建
	PostPermissionMethod PermissionMethod = "POST"
	// PutPermissionMethod 更新
	PutPermissionMethod PermissionMethod = "PUT"
	// DeletePermissionMethod 删除
	DeletePermissionMethod PermissionMethod = "DELETE"
	// PatchPermissionMethod 更新
	PatchPermissionMethod PermissionMethod = "PATCH"
	// OptionsPermissionMethod 选项
	OptionsPermissionMethod PermissionMethod = "OPTIONS"
	// HeadPermissionMethod 头部
	HeadPermissionMethod PermissionMethod = "HEAD"
)

// PermissionMethods 权限方法常量列表
var PermissionMethods = []PermissionMethod{GetPermissionMethod, PostPermissionMethod, PutPermissionMethod, DeletePermissionMethod, PatchPermissionMethod, OptionsPermissionMethod, HeadPermissionMethod}

// String 转换为字符串
func (p PermissionMethod) String() string {
	return string(p)
}

// Contains 是否包含
func (p PermissionMethod) Contains(method PermissionMethod) bool {
	for _, v := range PermissionMethods {
		if v == method {
			return true
		}
	}
	return false
}

// ContainsString 是否包含
func (p PermissionMethod) ContainsString(method string) bool {
	for _, v := range PermissionMethods {
		if v.String() == method {
			return true
		}
	}
	return false
}
