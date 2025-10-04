package main

import "fmt"

func main() {
	login:=false

	// normal if-else
	if login {
		fmt.Println("User is logged in")
	} else {
		fmt.Println("Please login!!")
	}

	// if else if ladder
	age:=35

	if age>40{
		fmt.Println("Person is in 40s")
	} else if age>30 && age<40 {
		fmt.Println("Person is in 30s")
	} else if age > 20{
		fmt.Println("Person is in 20s")
	} else{
		fmt.Println("Person is minor")
	}

	// using logical operators
	role:="admin"
	hasPermissions:=true

	if role=="admin" && hasPermissions {
		fmt.Println("Person is admin")
	} else {
		fmt.Println("Person is not an admin")
	}

	// declare directly iside if
	if age1:=12; age1>=18 {
		fmt.Println("Person is eligible")
	} else {
		fmt.Println("Person is teenage and not eligile")
	}

	// go doesn't have ternary operator
	
}
