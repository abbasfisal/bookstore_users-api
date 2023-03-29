package main

import (
	"bookstore_users-api/app"
	_ "bookstore_users-api/databases/mysql"
)

func main() {

	app.StartApplication()
}
