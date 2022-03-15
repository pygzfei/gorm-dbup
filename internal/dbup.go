package plugins

import (
	"fmt"
	"gorm-dbup/internal/utils"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
	"time"
)

type migrationUpdate struct {
	Version   string    `json:"version" gorm:"type:varchar(255);not null;"`
	ApplyTime time.Time `json:"apply_time" gorm:"type:datetime; not NULL"`
}

type IHandleMigration interface {
	GetVersion() string
	Handle(db *gorm.DB) (err error)
}

type DBUp struct {
	Database   string
	SQLFileDir string
}

func (m DBUp) Name() string {
	return "MysqlDbUp"
}

func (m DBUp) Initialize(db *gorm.DB) error {
	err := m.initRecordTable(db)
	if err != nil {
		return err
	}
	return m.doMigration(db)
}

func (m *DBUp) initRecordTable(db *gorm.DB) (err error) {
	var isExist int64

	db.Raw(fmt.Sprintf("SELECT COUNT(*) from information_schema.TABLES WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = 'migration_update'", m.Database)).Scan(&isExist)

	if isExist > 0 {
		return nil
	} else {
		// 没有运行过
		tableName := "CREATE TABLE `migration_update` " +
			"(`version` varchar(255) NOT NULL,`apply_time` datetime NOT NULL,PRIMARY KEY (`version`)) " +
			"ENGINE=InnoDB DEFAULT CHARSET=utf8;"
		exec := db.Exec(tableName)
		if exec.Error != nil {
			return exec.Error
		}
	}
	return err
}

func (m *DBUp) doMigration(db *gorm.DB) (err error) {
	names := &[]string{}
	err = db.Table("migration_update").Select("version").Find(names).Error
	if err != nil {
		return err
	}
	return db.Transaction(func(tx *gorm.DB) error {
		filePaths, err := readSqlFile(m.SQLFileDir)
		if err != nil {
			return err
		}
		for _, path := range filePaths {
			if isContainer(names, path) {
				continue
			}
			readFile, err := ioutil.ReadFile(m.SQLFileDir + "/" + path)
			if err != nil {
				return err
			}
			requests := strings.Split(string(readFile), ";")
			for _, request := range requests {
				q := strings.TrimSpace(request)
				if q == "" {
					continue
				}
				all := strings.ReplaceAll(q, `'`, `"`)
				sqlString := utils.TrimSqlString(all)
				err = db.Exec(sqlString).Error
				if err != nil {
					return err
				}
			}
			err = tx.Create(&migrationUpdate{Version: path, ApplyTime: time.UnixMilli(time.Now().UnixMilli())}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func isContainer(arr *[]string, str string) bool {
	for _, s := range *arr {
		if s == str {
			return true
		}
	}
	return false
}

func readSqlFile(sqlFileDir string) (filePaths []string, err error) {
	fileInfos, err := ioutil.ReadDir(sqlFileDir)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
		if strings.HasSuffix(strings.ToLower(fileInfo.Name()), ".sql") {
			filePaths = append(filePaths, fileInfo.Name())
		}
	}
	return filePaths, err
}
