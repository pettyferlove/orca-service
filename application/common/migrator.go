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
	)
	if err != nil {
		return err
	}
	return nil
}
