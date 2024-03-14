package common

import (
	"context"
	"github.com/google/uuid"
	bm "orca-service/application/model"
	"orca-service/global"
	"orca-service/global/config"
	"orca-service/global/security"
	"orca-service/global/security/model"
	"orca-service/global/util"
	"testing"
	"time"
)

func TestConnectMysql(t *testing.T) {
	host := "localhost"
	port := "3306"
	username := "root"
	password := "123456"
	database := "orca_test"

	_, err := ConnectMysql(host, port, username, password, database)
	if err != nil {
		t.Errorf("Failed to connect to mysql: %v", err)
	}
}

func TestInitDatabase(t *testing.T) {
	global.Config.Database = config.Database{
		Host:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		Database: "orca_test",
		Pool: config.DatabaseConnectionPool{
			MaxOpenConnection: 100,
			MaxIdleConnection: 10,
			MaxLifetime:       30,
			IdleTimeout:       30,
		},
	}

	err := InitDatabase()
	if err != nil {
		t.Errorf("Failed to initialize database: %v", err)
	}
}

func TestCreate(t *testing.T) {
	// 创建一个context
	ctx := context.Background()
	global.Config.Database = config.Database{
		Host:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		Database: "orca_test",
		Pool: config.DatabaseConnectionPool{
			MaxOpenConnection: 100,
			MaxIdleConnection: 10,
			MaxLifetime:       30,
			IdleTimeout:       30,
		},
	}
	err := InitDatabase()
	if err != nil {
		t.Errorf("Failed to initialize database: %v", err)
	}

	ctx = util.WithContext(ctx, &security.JWTClaims{
		UserDetail: model.UserDetail{
			Id: "0000001",
		},
		Roles: []string{
			"ROLE_ADMIN",
			"ROLE_USER",
		},
		Permissions: []string{
			"sys:user:*",
			"api:hello:get",
		},
	})
	err = global.DatabaseClient.Migrator().AutoMigrate(&bm.User{})
	if err != nil {
		t.Errorf("Failed to migrate database: %v", err)
		return
	}

	db := global.DatabaseClient.WithContext(ctx)
	test1 := uuid.New().String()
	db.Create(&bm.User{
		Id:                test1,
		LoginName:         "test_1",
		Password:          "123456",
		TenantId:          "0000001",
		Channel:           "test",
		Status:            1,
		LoginFail:         0,
		LastLoginFailTime: time.Now(),
	})
	// 查询
	var user bm.User
	db.Where("id = ?", test1).First(&user)
	if user.Id != test1 {
		t.Errorf("Failed to create user: %v", user)
	} else {
		t.Logf("Create user successfully: %v", user)
	}

	// 批量创建用户
	var users []bm.User
	for i := 0; i < 10; i++ {
		users = append(users, bm.User{
			Id:                uuid.New().String(),
			LoginName:         "test_" + string(rune(i)),
			Password:          "123456",
			TenantId:          "0000001",
			Channel:           "test",
			Status:            1,
			LoginFail:         0,
			LastLoginFailTime: time.Now(),
		})
	}
	db.Create(&users)
	// 查询总数
	var count int64
	db.Model(&bm.User{}).Count(&count)
	if count != 11 { // 11 = 1 + 10
		t.Errorf("Failed to batch create users: %v", count)
	} else {
		t.Logf("Batch create users successfully: %v", count)
	}
	err = global.DatabaseClient.Migrator().DropTable(&bm.User{})
	if err != nil {
		t.Errorf("Failed to drop table: %v", err)
		return
	}
}
