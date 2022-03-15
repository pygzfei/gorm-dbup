package test

import (
	"gorm-dbup/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestCanRun(t *testing.T) {
	database, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN:                       "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}))

	if err != nil {
		t.Error(err)
	}
	migration := pkg.NewMigration("test", "./dbup")
	err = database.Use(migration)

	if err != nil {
		t.Error(err)
	}
}
