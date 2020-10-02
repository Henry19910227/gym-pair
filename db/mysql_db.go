package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db       *sql.DB
	username string = "henry"
	pwd      string = "aaaa8027"
	address  string = "127.0.0.1:8889"
	database string = "ichat"
)

// NewDB 取得DB
func NewDB() *sql.DB {
	if db == nil {
		var err error
		datasource := fmt.Sprintf("%v:%v@tcp(%v)/%v", username, pwd, address, database)
		db, err = sql.Open("mysql", datasource)
		if err != nil {
			log.Fatal(err)
		}
	}
	return db
}
