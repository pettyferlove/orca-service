package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"orca-service/application/entity"
	"orca-service/global"
	"orca-service/global/handler"
	"orca-service/global/logger"
	"orca-service/global/security"
	"orca-service/global/security/model"
	"orca-service/global/security/token"
)

type Token struct {
	handler.Handler
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Type  string `json:"type,default=Bearer"`
}

func (t Token) Create(c *gin.Context) {
	var loginRequest LoginRequest
	err := t.MakeContext(c).Bind(&loginRequest).Errors
	if err != nil {
		logger.Error(err.Error())
		t.ResponseInvalidArgument(err.Error())
		return
	}
	var userEntity = entity.User{}
	tx := t.DataBase.Where("username = ?", loginRequest.Username).First(&userEntity)
	if tx.Error != nil {
		t.ResponseUnauthorizedMessage(tx.Error.Error())
		return
	}
	if userEntity.Id == "" {
		t.ResponseUnauthorizedMessage("User not found")
		return
	}
	var userInfoEntity = entity.UserInfo{}
	tx = t.DataBase.Where("user_id = ?", userEntity.Id).First(&userInfoEntity)
	if tx.Error != nil {
		t.ResponseUnauthorizedMessage(tx.Error.Error())
		return
	}
	if userInfoEntity.Id == "" {
		t.ResponseUnauthorizedMessage("User not found")
		return
	}
	// 查询角色列表
	var roleEntities []entity.Role
	tx = t.DataBase.Select("id", "role").Where("exists (select 1 from s_user_role where s_user_role.role_id = s_role.id and s_user_role.user_id = ?)", userEntity.Id).Find(&roleEntities)
	if tx.Error != nil {
		t.ResponseUnauthorizedMessage(tx.Error.Error())
		return
	}
	var roles []string
	var roleIds []string
	for _, role := range roleEntities {
		roleIds = append(roleIds, role.Id)
		roles = append(roles, role.Role)
	}
	// 查询角色权限
	var permissionEntities []entity.Permission
	t.DataBase.Select("permission").Where("exists (select 1 from s_role_permission where s_role_permission.permission_id = s_permission.id and s_role_permission.role_id in ?)", roleIds).Find(&permissionEntities)

	var permissions []string
	for _, permission := range permissionEntities {
		permissions = append(permissions, permission.Permission)
	}

	detail := model.UserDetail{
		Id:          userEntity.Id,
		Username:    userEntity.Username,
		Password:    userEntity.Password,
		Avatar:      userInfoEntity.Avatar,
		Nickname:    userInfoEntity.Nickname,
		Email:       userInfoEntity.Email,
		Phone:       userInfoEntity.Phone,
		Channel:     userEntity.Channel,
		Tenant:      userEntity.TenantId,
		Status:      userEntity.Status,
		Roles:       roles,
		Permissions: permissions,
	}

	// 使用BCrypt进行密码校验
	err = bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(loginRequest.Password))
	if err != nil {
		t.ResponseUnauthorizedMessage("Password error")
		return
	}
	store := token.NewRedisStore(global.RedisClient)
	accessToken, err := store.CreateAccessToken(detail)
	if err != nil {
		return

	}
	if err != nil {
		t.ResponseUnauthorizedMessage("Token creation failed")
		return
	}
	t.Response(LoginResponse{Token: accessToken, Type: "Bearer"})
	return
}

func (t Token) Delete(c *gin.Context) {
	t.MakeContext(c)
	// JWT Token无需删除，客户端扔掉即可，因为它是短期的，服务端不需记录它
	// 兼容后期Token加入Redis或者JWT Token加入Redis
	t.ResponseOk()
	return
}

// Refresh 刷新Token
func (t Token) Refresh(context *gin.Context) {
	t.MakeContext(context)
	originalClaims, exists := t.Context.Get("original_claims")
	if !exists {
		t.ResponseBusinessError(1, "Token parsing failed")
		return
	} else {
		// 转换为JWTClaims
		var _ = originalClaims.(*security.JWTClaims)

		t.Response("")
		return
	}
}
