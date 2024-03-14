package common

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"orca-service/global"
	"testing"
)

func TestMigrate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		DefaultStringSize:         256,  // default size for string fields
		DisableDatetimePrecision:  true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
	}

	global.DatabaseClient = gormDB
	mock.ExpectExec("CREATE TABLE").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("CREATE TABLE").WillReturnResult(sqlmock.NewResult(1, 1))
	err = Migrate()
	if err != nil {
		t.Errorf("Failed to migrate: %v", err)
	}
}
