package main

import "fmt"


func main()  {
	a:=10
	b:=15

	a=a+b

	if b>=a{
		b=b-a
	} else {
		b=a-b
	}

	a=a-b

	fmt.Println(a, b)

	fmt.Print("Using bitwise operator XOR!! ")
	fmt.Println()

	x:=11
	y:=13

	x=x^y
	fmt.Println(x, y)
	y=x^y
	fmt.Println(x, y)
	x=x^y

	fmt.Println(x, y)

}