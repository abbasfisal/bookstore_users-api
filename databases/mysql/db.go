package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Db *sql.DB
)

func init() {

	var err error
	Db, err = sql.Open("mysql", "root:root@/golang")

	if err != nil {
		panic(err.Error())
	}
	pErr := Db.Ping()
	if pErr != nil {
		log.Fatal(pErr.Error())
	}
	fmt.Println("ping database successfully")

	fmt.Println("Connected to db successfully")
}
