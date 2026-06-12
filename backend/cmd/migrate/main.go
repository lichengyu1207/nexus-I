package main

import (
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 默认数据库路径
	dbPath := "./data/dev.db"
	if len(os.Args) > 1 {
		dbPath = os.Args[1]
	}

	// 确保 data 目录存在
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("创建目录失败: %v", err)
	}

	// 连接数据库
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("打开数据库失败: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 执行 migrations 目录下的所有 .sql 文件
	migrationsDir := "migrations"
	entries, err := fs.ReadDir(os.DirFS(migrationsDir), ".")
	if err != nil {
		log.Fatalf("读取迁移目录失败: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".sql" {
			continue
		}

		sqlFile := filepath.Join(migrationsDir, entry.Name())
		content, err := os.ReadFile(sqlFile)
		if err != nil {
			log.Printf("读取文件 %s 失败: %v", entry.Name(), err)
			continue
		}

		if _, err := db.Exec(string(content)); err != nil {
			log.Printf("执行迁移 %s 失败: %v", entry.Name(), err)
			continue
		}

		fmt.Printf("✅ 迁移完成: %s\n", entry.Name())
	}

	fmt.Println("\n数据库初始化完毕:", dbPath)
}
