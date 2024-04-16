package model

import "orca-service/application/constant"

// PermissionCreate 权限创建
type PermissionCreate struct {
	ParentID           string                    `json:"parentId" validate:"omitempty,uuid4"`
	Permission         string                    `json:"permission" validate:"required"`
	PermissionName     string                    `json:"permissionName"`
	PermissionType     constant.PermissionType   `json:"permissionType"`
	PermissionMetadata string                    `json:"permissionMetadata"`
	PermissionMethod   constant.PermissionMethod `json:"permissionMethod"`
	PermissionURL      string                    `json:"permissionURL"`
	Remark             string                    `json:"remark"`
	Valid              bool                      `json:"valid"`
	Visible            bool                      `json:"visible"`
	Icon               string                    `json:"icon"`
}

// PermissionUpdate 权限更新
type PermissionUpdate struct {
}

// PermissionDetail 权限详情
type PermissionDetail struct {
}

// PermissionTreeRequest 权限树请求参数
type PermissionTreeRequest struct {
}

// PermissionTree 权限树
type PermissionTree struct {
}

// PermissionAvailableTree 权限可用树
type PermissionAvailableTree struct {
}
