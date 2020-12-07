package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func init() {
	var err error
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "pass.1234",
		"127.0.0.1", "3306", "testdb")
	SqlDB, err = sql.Open("mysql", conn)
	if err != nil {
		log.Fatalf("connection mysql fail:%v", err)
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatalf("dail mysql fail:%v", err)
	}
}