package geeorm

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSqlite(t *testing.T) {
	db, _ := sql.Open("sqlite3", "gee.db")
	defer func() { _ = db.Close() }()
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")
	result, err := db.Exec("INSERT INTO User(`NAME`) values (?),(?)", "Tom", "Sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}
	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
}

func TestTransaction(t *testing.T) {
	db, _ := sql.Open("sqlite3", "gee.db")
	defer func() { _ = db.Close() }()
	_, _ = db.Exec("CREATE TABLE IF NOT EXISTS User('Name' text)")

	tx, _ := db.Begin()
	_, err1 := tx.Exec("INSERT INTO User(`Name`) VALUES (?)", "Tom")
	_, err2 := tx.Exec("INSERT INTO User(`Name`) VALUES (?)", "Jack")
	if err1 != nil || err2 != nil {
		_ = tx.Rollback()
		log.Println("Rollback", err1, err2)
	} else {
		_ = tx.Commit()
		log.Println("Commit")
	}

}
