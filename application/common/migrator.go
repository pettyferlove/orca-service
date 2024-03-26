package common

import (
	"orca-service/application/entity"
	"orca-service/global"
)

func Migrate() error {
	migrator := global.DatabaseClient.Migrator()
	err := migrator.AutoMigrate(
		&entity.User{},
		&entity.UserInfo{},
		&entity.Role{},
		&entity.Permission{},
		&entity.Menu{},
		&entity.UserRole{},
		&entity.RoleMenu{},
		&entity.RolePermission{},
	)
	if err != nil {
		return err
	}
	return nil
}
