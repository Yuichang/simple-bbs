package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// SQLに接続する。
func ConnectDB() (*sql.DB, error) {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// DSNを構築
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_ADDRESS"),
		DBName: os.Getenv("DB_NAME"),
	}
	// DBハンドルの作成
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	// DBへの接続確認
	if err := db.Ping(); err != nil {
		db.Close()
		fmt.Println("Ping error", err)
		return nil, err
	}
	fmt.Println("connection success!")
	return db, nil
}
