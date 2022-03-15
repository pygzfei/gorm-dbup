package pkg

import plugins "gorm-dbup/internal"

func NewMigration(database string, SQLFileDir string) *plugins.DBUp {
	return &plugins.DBUp{Database: database, SQLFileDir: SQLFileDir}
}
