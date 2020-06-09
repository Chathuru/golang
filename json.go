package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Phone string
}

type Employe struct {
	Emp_no     int
	Birth_date string
	First_name string
	Gender     string
	Hire_date  string
	Last_name  string
}

func main() {
	p1 := Person{"Chathura", 26, "071"}
	p2 := Person{"Chathura", 26, "071"}
	p3 := Person{"Chathura", 26, "071"}

	plist := []Person{p1, p2, p3, {"Lochana", 27, "071"}}
	fmt.Println(plist)
	byte, _ := json.Marshal(plist)
	fmt.Println(string(byte))

	e1 := Employe{10001, "1953-09-02", "Georgi", "M", "1986-06-26", "Facello"}
	e2 := Employe{10002, "1964-06-02", "Bezalel", "F", "1985-11-21", "Simmel"}

	elist := []Employe{e1, e2}
	fmt.Println(elist)
	byteList, err := json.Marshal(elist)
	fmt.Println(err)
	fmt.Println(string(byteList))
}
