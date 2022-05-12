package pkg

import plugins "github.com/pygzfei/gorm-dbup/internal"

// NewMigration creates a new migration
func NewMigration(database string, SQLFileDir string) *plugins.DBUp {
	return &plugins.DBUp{Database: database, SQLFileDir: SQLFileDir}
}
