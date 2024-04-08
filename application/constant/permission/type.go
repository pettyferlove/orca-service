package permission

// Type 权限类型常量
type Type string

const (
	// PAGE 页面
	PAGE Type = "PAGE"
	// API 接口
	API Type = "API"
	// BUTTON 按钮
	BUTTON Type = "BUTTON"
	// FOLDER 文件夹
	FOLDER Type = "FOLDER"
)

// String 转换为字符串
func (p Type) String() string {
	return string(p)
}
