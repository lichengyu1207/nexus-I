package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Iichengyu1207/nexus-I/backend/internal/config"
)

// NewDB 根据配置创建数据库连接
func NewDB(cfg *config.Config) (*sql.DB, error) {
	if cfg.DBType == "sqlite" {
		db, err := sql.Open("sqlite3", cfg.SQLitePath)
		if err != nil {
			return nil, err
		}
		if err := db.Ping(); err != nil {
			return nil, err
		}
		return db, nil
	}
	// PostgreSQL 分支稍后添加
	return nil, nil
}
