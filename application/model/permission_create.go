package model

type PermissionCreate struct {
	ParentID           string `json:"parentId"`
	Permission         string `json:"permission"`
	PermissionName     string `json:"permissionName"`
	PermissionType     string `json:"permissionType"`
	PermissionMetadata string `json:"permissionMetadata"`
	PermissionMethod   string `json:"permissionMethod"`
	PermissionURL      string `json:"permissionURL"`
	Remark             string `json:"remark"`
	Valid              bool   `json:"valid"`
	Visible            bool   `json:"visible"`
	Icon               string `json:"icon"`
}
