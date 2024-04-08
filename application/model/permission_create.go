package model

import "orca-service/application/constant"

type PermissionCreate struct {
	ParentID           string                    `json:"parentId" validate:"omitempty,uuid4"`
	Permission         string                    `json:"permission"`
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
