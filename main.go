package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

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
	_, err = Db.Exec(cmd, "hanako", 19)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = Db.Exec(cmd, "abe", 78)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\n==================================================================\n")

	// UPDATE
	cmd = "UPDATE persons SET age = ? WHERE name = ?"
	_, err = Db.Exec(cmd, 30, "tarou")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\n==================================================================\n")

	// SELECT
	cmd = "SELECT * FROM persons WHERE name = ?"
	row := Db.QueryRow(cmd, "tarou") // `QueryRow` は1レコードのみ取得
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No Row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println("p.Name, p.Age =>", p.Name, p.Age)
	fmt.Println("\n==================================================================\n")

	// SELECT All
	cmd = "SELECT * FROM persons"
	rows, _ := Db.Query(cmd) // クエリ結果を返す。今回は条件にマッチしたレコードを全て返す。
	defer rows.Close()

	// personsというスライスを作って、データを1個1個追加していってる。
	var persons []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil{
			log.Println(err)
		}
		persons = append(persons, p)
	}
	fmt.Println("persons =>", persons)

	for i, person := range persons {
		fmt.Println("i, person =>", i, person)
	}
	fmt.Println("\n==================================================================\n")


	// DELETE
	cmd = "DELETE FROM persons WHERE name = ?"
	_, err = Db.Exec(cmd, "tarou")
	if err != nil {
		log.Fatalln(err)
	}

	// DELETEできたか確認
	cmd = "SELECT * FROM persons"
	rows, _ = Db.Query(cmd) // クエリ結果を返す。今回は条件にマッチしたレコードを全て返す。
	defer rows.Close()

	persons = []Person{}
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil{
			log.Println(err)
		}
		persons = append(persons, p)
	}
	fmt.Println("persons =>", persons)
	fmt.Println("\n==================================================================\n")

}
