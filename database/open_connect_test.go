package database

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetConnection(t *testing.T) {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/comment")
	defer db.Close()

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

}
