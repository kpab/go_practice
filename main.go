package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// database + sqlite3
// テーブル作成
var Db *sql.DB

func main () {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	// cmd := `CREATE TABLE IF NOT EXISTS persons(
	// 			name STRING,
	// 			age INT)`
	
	// データの追加
	cmd := `INSERT INTO persons (name, age) VALUES (?, ?)`
	_, err := Db.Exec(cmd, "tarou", 20)
	if err != nil {
		log.Fatalln(err)
	}

	// データの更新：省略

}