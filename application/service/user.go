package service

import (
	"orca-service/application/constant"
	"orca-service/application/entity"
	"orca-service/application/model"
	"orca-service/global"
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
	// 优化子查询

	if err := u.DataBase.
		Model(&entity.User{}).
		Preload("UserInfo").
		Preload("Roles").
		Preload("Roles.Permissions", "permission_type in (?)", []string{constant.API.String(), constant.BUTTON.String()}).
		Where("username = ?", username).First(&userEntity).Error; err != nil {
		u.AddError(err)
		return nil
	}
	var roles = make([]string, 0)
	var permissions = make([]string, 0)
	for _, rolesEntity := range userEntity.Roles {
		roles = append(roles, rolesEntity.Role)
		for _, permission := range rolesEntity.Permissions {
			permissions = append(permissions, permission.Permission)
		}
	}
	detail := user.UserDetail{
		Id:          userEntity.Id,
		Username:    userEntity.Username,
		Password:    userEntity.Password,
		Avatar:      userEntity.UserInfo.Avatar,
		Nickname:    userEntity.UserInfo.Nickname,
		Email:       userEntity.UserInfo.Email,
		Phone:       userEntity.UserInfo.Phone,
		Channel:     userEntity.Channel,
		Status:      userEntity.Status,
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
	loginAttempts := model.LoginAttempts{}
	if err := u.DataBase.Model(&entity.User{}).Where("username = ?", username).First(&loginAttempts).Error; err != nil {
		u.AddError(err)
		return nil
	} else {
		return &loginAttempts
	}
}

// LoginFailed 方法用于处理登录失败后的逻辑
func (u *User) LoginFailed(username string) {
	tx := u.DataBase.Begin()
	security := global.Config.Security
	loginAttempts := u.LoadLoginAttempts(username)
	if loginAttempts == nil {
		return
	}
	if loginAttempts.LastLoginFailTime.IsZero() {
		err := u.LoginFailedForFailNum(username, 1)
		if err != nil {
			u.AddError(err)
		}
	} else {
		if time.Now().Sub(loginAttempts.LastLoginFailTime) > time.Duration(security.LoginAttempt.LockingDuration) {
			err := u.LoginFailedForFailNum(username, loginAttempts.LoginFail+1)
			if err != nil {
				u.AddError(err)
			}
		} else {
			err := u.LoginFailedForFailNum(username, loginAttempts.LoginFail+1)
			if err != nil {
				u.AddError(err)
			}
		}
	}
	if loginAttempts.LoginFail+1 >= security.LoginAttempt.TimesBeforeLock {
		err := u.LockUser(username)
		if err != nil {
			u.AddError(err)
		}
	}
	if u.Errors != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

// LoginFailedForFailNum 方法用于处理登录失败次数过多后的逻辑
func (u *User) LoginFailedForFailNum(username string, failNum int) error {
	if err := u.DataBase.Model(&entity.User{}).Where("username = ?", username).Updates(map[string]interface{}{"login_fail": failNum, "last_login_fail_time": time.Now()}).Error; err != nil {
		return err
	} else {
		return nil
	}
}

// LockUser 方法用于锁定用户
func (u *User) LockUser(username string) error {
	// 业务逻辑
	if err := u.DataBase.Model(&entity.User{}).Where("username = ?", username).Updates(map[string]interface{}{"status": user.Locked}).Error; err != nil {
		return err
	} else {
		return nil
	}
}
