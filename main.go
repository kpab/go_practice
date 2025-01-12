package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// database + sqlite3
// テーブル作成
var Db *sql.DB

var err error

type Person struct {
	name string
	age int
}

func main () {
	Db, err = sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	defer Db.Close()

	/*
	// C
	cmd := `INSERT INTO persons (name, age) VALUES ($1, $2) `
	_, err := Db.Exec(cmd, "Alex", 20)
	if err != nil {
		log.Fatalln(err)
	}
		*/
	/*
	// R
	cmd := `SELECT * FROM persons WHERE age = $1 `
	// QueryRow 1レコード取得
	row := Db.QueryRow(cmd, 20)
	var p Person
	err = row.Scan(&p.name, &p.age)
	if err != nil {
		// データが無かったら
		if err == sql.ErrNoRows {
			log.Println("no row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.name, p.age)

	cmd = `SELECT * FROM persons`
	// Queryは条件に合うものを全て取得
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	// structを作成
	var pp []Person
	// 取得したデータをループでスライスに追加 for rows.Next
	for rows.Next() {
		var p Person
		// scan データ追加
		err := rows.Scan(&p.name, &p.age)
		// 1行ずつエラーハンドリングver
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	// まとめてエラーハンドリングver
	err := rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	// 表示
	for _, p := range pp {
		fmt.Println(p.name, p.age)
	}
	*/

	// U
	cmd := `UPDATE persons SET age = $1 WHERE name = $2`
	_, err := Db.Query(cmd, 22, "Alex")
	if err != nil {
		log.Fatalln(err)
	}

	// D 省略


}