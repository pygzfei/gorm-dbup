package plugins

import (
	"errors"
	"fmt"
	"github.com/pygzfei/gorm-dbup/internal/utils"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
	"time"
)

// migrationUpdate is the migration update struct
type migrationUpdate struct {
	Version   string    `json:"version" gorm:"type:varchar(255);not null;"`
	ApplyTime time.Time `json:"apply_time" gorm:"type:datetime; not NULL"`
}

// IHandleMigration is the interface for migration
type IHandleMigration interface {
	GetVersion() string
	Handle(db *gorm.DB) (err error)
}

// DBUp is the struct for dbup
type DBUp struct {
	Database   string
	SQLFileDir string
}

// Name is the name of dbup
func (m DBUp) Name() string {
	return "MysqlDbUp"
}

// Initialize is the init function of dbup
func (m DBUp) Initialize(db *gorm.DB) error {
	err := m.initRecordTable(db)
	if err != nil {
		return err
	}
	return m.doMigration(db)
}

// initRecordTable is the function to init record table
func (m *DBUp) initRecordTable(db *gorm.DB) (err error) {
	var databaseIsExist int64
	var tableIsExist int64

	db.Raw(fmt.Sprintf("SELECT COUNT(information_schema.SCHEMATA.SCHEMA_NAME) FROM information_schema.SCHEMATA where SCHEMA_NAME='%s'", m.Database)).Scan(&databaseIsExist) // check database

	if databaseIsExist <= 0 {
		if err != nil {
			return errors.New(fmt.Sprintf("%s not exist", m.Database))
		}
	}

	db.Raw(fmt.Sprintf("SELECT COUNT(*) from information_schema.TABLES WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = 'migration_update'", m.Database)).Scan(&tableIsExist)

	if tableIsExist > 0 {
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

// doMigration is the function to do migration
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
			if len(readFile) <= 0 {
				continue
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
			err = tx.Exec(fmt.Sprintf("INSERT INTO migration_update (version, apply_time) VALUES ('%s', '%s')", path, time.Now().Format("2006-01-02 15:04:05"))).Error
			if err != nil {
				return err
			}
			fmt.Println(fmt.Sprintf("Migration %s success", path))
		}
		return nil
	})
}

// isContainer is the function to check if the file is container
func isContainer(arr *[]string, str string) bool {
	for _, s := range *arr {
		if s == str {
			return true
		}
	}
	return false
}

// readSqlFile is the function to read sql file
func readSqlFile(sqlFileDir string) (filePaths []string, err error) {
	fileInfos, err := ioutil.ReadDir(sqlFileDir)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		if strings.HasSuffix(strings.ToLower(fileInfo.Name()), ".sql") {
			filePaths = append(filePaths, fileInfo.Name())
		}
	}
	return filePaths, err
}
