package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// InitSQLite 初始化 SQLite 连接
func InitSQLite(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
