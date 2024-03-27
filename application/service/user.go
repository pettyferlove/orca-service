package service

import (
	"orca-service/application/entity"
	"orca-service/application/model"
	user "orca-service/global/security"
	"orca-service/global/service"
	"time"
)

type User struct {
	service.Service
}

// LoadUserByUsername 方法用于根据用户名加载用户信息
func (u *User) LoadUserByUsername(username string) *user.UserDetail {
	var userEntity = entity.User{}
	if err := u.DataBase.Where("username = ?", username).First(&userEntity).Error; err != nil {
		u.AddError(err)
		return nil
	}
	var userInfoEntity = entity.UserInfo{}
	if err := u.DataBase.Where("user_id = ?", userEntity.Id).First(&userInfoEntity).Error; err != nil {
		u.AddError(err)
		return nil
	}
	// 查询角色列表
	var roleEntities []entity.Role
	if err := u.DataBase.Select("id", "role").Where("exists (select 1 from s_user_role where s_user_role.role_id = s_role.id and s_user_role.user_id = ?)", userEntity.Id).Find(&roleEntities).Error; err != nil {
		u.AddError(err)
		return nil
	}
	var roles []string
	var roleIds []string
	for _, role := range roleEntities {
		roleIds = append(roleIds, role.Id)
		roles = append(roles, role.Role)
	}
	// 查询角色权限
	var permissionEntities []entity.Permission
	if err := u.DataBase.Select("permission").Where("exists (select 1 from s_role_permission where s_role_permission.permission_id = s_permission.id and s_role_permission.role_id in ?)", roleIds).Find(&permissionEntities).Error; err != nil {
		u.AddError(err)
		return nil

	}
	var permissions []string
	for _, permission := range permissionEntities {
		permissions = append(permissions, permission.Permission)
	}
	detail := user.UserDetail{
		Id:          userEntity.Id,
		Username:    userEntity.Username,
		Password:    userEntity.Password,
		Avatar:      userInfoEntity.Avatar,
		Nickname:    userInfoEntity.Nickname,
		Email:       userInfoEntity.Email,
		Phone:       userInfoEntity.Phone,
		Channel:     userEntity.Channel,
		Tenant:      userEntity.TenantId,
		Status:      user.Normal,
		Roles:       roles,
		Permissions: permissions,
	}
	return &detail
}

// LoginSuccess 方法用于处理登录成功后的逻辑
func (u *User) LoginSuccess(username string) {
	if err := u.DataBase.Model(&entity.User{}).Where("username = ?", username).Updates(map[string]interface{}{"login_fail": 0, "last_login_fail_time": nil}).Error; err != nil {
		u.AddError(err)
		return
	}
}

// LoadLoginAttempts 方法用于加载登录尝试次数
func (u *User) LoadLoginAttempts(username string) *model.LoginAttempts {
	var userEntity = entity.User{}
	if err := u.DataBase.Select("id", "username", "login_fail", "last_login_fail_time").Where("username = ?", username).First(&userEntity).Error; err != nil {
		u.AddError(err)
		return nil
	} else {
		return &model.LoginAttempts{
			Id:                userEntity.Id,
			Username:          userEntity.Username,
			LoginFailNum:      userEntity.LoginFail,
			LastLoginFailTime: userEntity.LastLoginFailTime,
		}
	}
}

// LoginFailed 方法用于处理登录失败后的逻辑
func (u *User) LoginFailed(username string) {
	// 业务逻辑
}

// LoginFailedForFailNum 方法用于处理登录失败次数过多后的逻辑
func (u *User) LoginFailedForFailNum(username string, failNum int) {
	if err := u.DataBase.Model(&entity.User{}).Where("username = ?", username).Updates(map[string]interface{}{"login_fail": failNum, "last_login_fail_time": time.Now()}).Error; err != nil {
		u.AddError(err)
		return
	}
}

// LockUser 方法用于锁定用户
func (u *User) LockUser(username string) {
	// 业务逻辑
	if err := u.DataBase.Model(&entity.User{}).Where("username = ?", username).Updates(map[string]interface{}{"status": user.Locked}).Error; err != nil {
		u.AddError(err)
		return
	}
}

// CheckLoginAttempts 方法用于检查登录尝试次数
func (u *User) CheckLoginAttempts(username string) {
	// 业务逻辑
}
