package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"orca-service/global"
	log "orca-service/global/logger"
	"orca-service/global/util"
	"reflect"
	"time"
)

// InitDatabase 初始化数据库连接并设置连接池参数
func InitDatabase() error {
	// 获取数据库配置
	databaseConfig := global.Config.Database
	// 连接到MySQL数据库
	connect, err := ConnectMysql(databaseConfig.Host, databaseConfig.Port, databaseConfig.Username, databaseConfig.Password, databaseConfig.Database)
	if err != nil {
		return err
	}
	// 获取数据库连接池配置
	poolConfig := databaseConfig.Pool
	// 获取数据库连接
	db, err := connect.DB()
	if err != nil {
		log.Error("database connection failed: %v", err)
		return err
	}
	// 设置数据库连接池参数
	db.SetMaxOpenConns(poolConfig.MaxOpenConnection)
	db.SetMaxIdleConns(poolConfig.MaxIdleConnection)
	db.SetConnMaxLifetime(time.Second * time.Duration(poolConfig.MaxLifetime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(poolConfig.IdleTimeout))

	err = connect.Callback().Create().Before("gorm:before_create").Register("create_audit", createAudit)
	if err != nil {
		log.Error("register create audit callback failed: %v", err)
		return err
	}
	err = connect.Callback().Update().Before("gorm:before_update").Register("update_audit", updateAudit)
	if err != nil {
		log.Error("register update audit callback failed: %v", err)
		return err

	}
	// 将数据库连接保存到全局变量中
	global.DatabaseClient = connect
	return nil
}

// ConnectMysql 连接到MySQL数据库
func ConnectMysql(host, port, username, password, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, database)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			// 设置日志级别
			log.DefaultLogger(),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info, // GORM 日志级别
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				Colorful:                  false,
			},
		),
	})
}

func createAudit(db *gorm.DB) {
	accountId := util.GetAccountId(db.Statement.Context)
	data := db.Statement.ReflectValue
	if data.Kind() == reflect.Slice {
		for i := 0; i < data.Len(); i++ {
			elem := data.Index(i)
			field := elem.FieldByName("Creator")
			if field.IsValid() && field.CanSet() {
				field.SetString(accountId)
			}
			field = elem.FieldByName("CreateTime")
			if field.IsValid() && field.CanSet() {
				field.Set(reflect.ValueOf(time.Now()))
			}
			field = elem.FieldByName("Modifier")
			if field.IsValid() && field.CanSet() {
				field.SetString(accountId)
			}
			field = elem.FieldByName("ModifyTime")
			if field.IsValid() && field.CanSet() {
				field.Set(reflect.ValueOf(time.Now()))
			}
		}
		return
	}
	if db.Statement.Schema.LookUpField("Creator") != nil {
		db.Statement.SetColumn("Creator", accountId, true)
	}
	if db.Statement.Schema.LookUpField("CreateTime") != nil {
		db.Statement.SetColumn("CreateTime", time.Now(), true)
	}
	if db.Statement.Schema.LookUpField("Modifier") != nil {
		db.Statement.SetColumn("Modifier", accountId, true)
	}
	if db.Statement.Schema.LookUpField("ModifyTime") != nil {
		db.Statement.SetColumn("ModifyTime", time.Now(), true)
	}
}

func updateAudit(db *gorm.DB) {
	accountId := util.GetAccountId(db.Statement.Context)
	data := db.Statement.ReflectValue
	if data.Kind() == reflect.Slice {
		for i := 0; i < data.Len(); i++ {
			elem := data.Index(i)
			field := elem.FieldByName("Modifier")
			if field.IsValid() && field.CanSet() {
				field.SetString(accountId)
			}
			field = elem.FieldByName("ModifyTime")
			if field.IsValid() && field.CanSet() {
				field.Set(reflect.ValueOf(time.Now()))
			}
		}
		return
	}
	if db.Statement.Schema.LookUpField("Modifier") != nil {
		db.Statement.SetColumn("Modifier", accountId, true)
	}
	if db.Statement.Schema.LookUpField("ModifyTime") != nil {
		db.Statement.SetColumn("ModifyTime", time.Now(), true)
	}
}
