package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	// DB作成
	cmd := `CREATE TABLE IF NOT EXISTS persons(
		name STRING,
		age INT)`
	_, err := Db.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// INSERT
	// INSERTされたか確認するには `sqlite3 example.sql` を実行して `select * from persons;` で分かる。
	cmd = "INSERT INTO persons (name, age) VALUES(?, ?)"
	_, err = Db.Exec(cmd, "tarou", 25)
	if err != nil {
		log.Fatalln(err)
	}

}
