package constant

// PermissionMethod 权限方法常量
type PermissionMethod string

// PermissionType 权限类型常量
type PermissionType string

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

const (
	// PagePermissionType 页面
	PagePermissionType PermissionType = "PAGE"
	// ApiPermissionType 接口
	ApiPermissionType PermissionType = "API"
	// ButtonPermissionType 按钮
	ButtonPermissionType PermissionType = "BUTTON"
	// FolderPermissionType 文件夹
	FolderPermissionType PermissionType = "FOLDER"
)

// PermissionMethods 权限方法常量列表
var PermissionMethods = []PermissionMethod{GetPermissionMethod, PostPermissionMethod, PutPermissionMethod, DeletePermissionMethod, PatchPermissionMethod, OptionsPermissionMethod, HeadPermissionMethod}

// PermissionTypeTypeList 权限类型列表
var PermissionTypeTypeList = []PermissionType{PagePermissionType, ApiPermissionType, ButtonPermissionType, FolderPermissionType}

// String 转换为字符串
func (p PermissionMethod) String() string {
	return string(p)
}

// String 转换为字符串
func (p PermissionType) String() string {
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
