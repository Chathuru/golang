package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(10.22.132.203)/employee")
	if err != nil {
		log.Fatalf("Connection error (%s)", err)
	}

	return db
}

func listAllEmployes(db *sql.DB) {
	result, _ := db.Query("SELECT * FROM employees LIMIT 20")
	employees := []Employe{}

	for result.Next() {
		var birth_date, first_name, gender, hire_date, last_name string
		var emp_no int
		result.Scan(&emp_no, &birth_date, &first_name, &last_name, &gender, &hire_date)
		employees = append(employees, Employe{emp_no, birth_date, first_name, gender, hire_date, last_name})
	}

	byte, _ := json.Marshal(employees)
	fmt.Println(string(byte))
}

func describeTable(db *sql.DB) {
	result, _ := db.Query("SELECT * FROM employees LIMIT 1")
	colum_names, _ := result.Columns()
	fmt.Println(colum_names)
}

func selectById(db *sql.DB, empNo int) {
	query_result := db.QueryRow("SELECT * FROM employees WHERE emp_no = ?", empNo)

	var name Employe
	query_result.Scan(&name.Emp_no, &name.Birth_date, &name.First_name, &name.Last_name, &name.Gender, &name.Hire_date)

	bytePerson, _ := json.Marshal(name)
	fmt.Println(string(bytePerson))
}

func insertEmploye(db *sql.DB, name string) {
	sqlInsert, err := db.Prepare("INSERT INTO gotable(Name) VALUES(?)")
	if err != nil {
		log.Fatalln(err)
	}

	result, err := sqlInsert.Exec(name)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalf("Insert success... (%s)", result)
}
