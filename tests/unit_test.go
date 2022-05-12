package tests

import (
	"github.com/pygzfei/gorm-dbup/pkg"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestCanRun(t *testing.T) {
	database, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN:                       "root:123456@tcp(localhost:3306)/test1?charset=utf8mb4&parseTime=True&loc=Local",
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}))

	if err != nil {
		t.Error(err)
	}

	if database.Exec("DROP TABLE IF EXISTS `migration_update`").Error != nil {
		t.Error(err)
	}

	if database.Use(pkg.NewMigration("test1", "./dbup1")) != nil {
		t.Error(err)
	}

	var migrationCount int
	if database.Raw("select Count(*) from migration_update").Scan(&migrationCount).Error != nil {
		t.Error(err)
	}

	var roleCount int
	if database.Raw("select Count(*) from role").Scan(&roleCount).Error != nil {
		t.Error(err)
	}

	var userCount int
	if database.Raw("select Count(*) from user").Scan(&userCount).Error != nil {
		t.Error(err)
	}
	assert.Equal(t, migrationCount, 2)
	assert.Equal(t, roleCount, 6)
	assert.Equal(t, userCount, 21)
}

func TestCanAddNewMigrationUpdate(t *testing.T) {
	database, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN:                       "root:123456@tcp(localhost:3306)/test1?charset=utf8mb4&parseTime=True&loc=Local",
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}))

	if err != nil {
		t.Error(err)
	}

	if database.Use(pkg.NewMigration("test1", "./dbup2")) != nil {
		t.Error(err)
	}

	var migrationCount int
	if database.Raw("select Count(*) from migration_update").Scan(&migrationCount).Error != nil {
		t.Error(err)
	}

	var roleCount int
	if database.Raw("select Count(*) from role").Scan(&roleCount).Error != nil {
		t.Error(err)
	}
	assert.Equal(t, roleCount, 7)
	assert.Equal(t, migrationCount, 3)
}
