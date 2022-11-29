package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	SqlBackend struct {
		*MysqlBackend
	}
)

func NewSqliteBackend(dsn string) (*SqlBackend, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn eg: "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &SqlBackend{
		MysqlBackend: &MysqlBackend{
			db: db,
		},
	}, nil
}
