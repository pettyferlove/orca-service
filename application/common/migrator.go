package common

import (
	"orca-service/application/model"
	"orca-service/global"
)

func Migrate() error {
	migrator := global.DatabaseClient.Migrator()
	err := migrator.AutoMigrate(
		&model.User{},
		&model.UserInfo{},
	)
	if err != nil {
		return err
	}
	return nil
}
