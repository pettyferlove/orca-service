package permission

// Method 权限方法常量
type Method string

const (
	// Get 获取
	Get = "GET"
	// Post 创建
	Post = "POST"
	// Put 更新
	Put = "PUT"
	// Delete 删除
	Delete = "DELETE"
	// Patch 更新
	Patch = "PATCH"
	// Options 选项
	Options = "OPTIONS"
	// Head 头部
	Head = "HEAD"
)

// Permissions 权限方法常量列表
var Permissions = []Method{Get, Post, Put, Delete, Patch, Options, Head}

// String 转换为字符串
func (p Method) String() string {
	return string(p)
}

// Contains 是否包含
func (p Method) Contains(m Method) bool {
	for _, v := range Permissions {
		if v == m {
			return true
		}
	}
	return false

}

// ContainsString 是否包含
func (p Method) ContainsString(m string) bool {
	for _, v := range Permissions {
		if v.String() == m {
			return true
		}
	}
	return false
}
