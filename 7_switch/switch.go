package main

import (
	"fmt"
	"reflect"
	"time"
)

func main()  {
	// simple switch
	i:=5

	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("None of the above")
	}

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Time for side projects and learning!")
	default:
		fmt.Println("Office time")
	}

	// type switch
	identify_type := func(i interface{}){
		switch t :=i.(type){
		case int:
			fmt.Println("Integer")
		case float32:
			fmt.Println("Float32")
		case float64:
			fmt.Println("Float64")
		case string:
			fmt.Println("String")
		default:
			fmt.Println("Dont know", reflect.TypeOf(t))
		
		}
	}

	identify_type(true)
}
