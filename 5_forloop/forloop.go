package main

import "fmt"

// for loop is the only loop in go

var x string = "holmes"

func main()  {
	// to execute while type loop in go we can do something like below

	i:=1

	for i<=5 {
		fmt.Println(i)
		i+=1
	}

	// infinite loop
	// for {
	// 	fmt.Println("yes")
	// }

	fmt.Println("Normal for loop printing...")
	// normal/classic for loop
	for i:=0; i<=5; i++ {
		fmt.Println(i)
		// we can even use break or continue also here in loop
		// break
		// continue
	}

	// in go 1.22 range function has been introduced

	for i:=range 3{
		// fmt.Println(i)
		fmt.Println(i, x)
	}


}


