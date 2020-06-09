package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := dbConnect()
	listAllEmployes(db)
	db.Close()
}
