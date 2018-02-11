package main

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
//go get github.com/go-sql-driver/mysql

func insert(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO go_test(id,username, password) VALUES(?,?, ?)")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return
	}
	stmt.Exec("1","guotie", "guotie")
	stmt.Exec("2","testuser", "123123")

}

func main() {
	//db, err := sql.Open("mysql", "mysql:Welcome1@/mysql")
	db, err := sql.Open("mysql", "mysql:Welcome1@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	insert(db)

	rows, err := db.Query("select id, username from go_test where id = ?", 1)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	var id int
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}