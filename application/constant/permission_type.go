package constant

// PermissionType 权限类型常量
type PermissionType string

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

// PermissionTypeTypeList 权限类型列表
var PermissionTypeTypeList = []PermissionType{PagePermissionType, ApiPermissionType, ButtonPermissionType, FolderPermissionType}

// String 转换为字符串
func (p PermissionType) String() string {
	return string(p)
}
