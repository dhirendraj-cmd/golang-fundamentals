package main

import "fmt"

var radius float64 = 2.0 // this is allowed

// radius:= not allowed

func main()  {

	// for declaring constants in go we use const keyword
	const pi=3.14

	var circle_area = pi*radius*radius

	fmt.Println("Area of circle is: ", circle_area)

	// constant grouping

	const (
		db_user="holmes"
		db_name="pidrive"
		port=8080
		host="locahost"
	)

	fmt.Println(db_user, db_name, port, host)

}

// Note a constant must be assigned value at the beginning



