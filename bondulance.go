package main

import "fmt"

var bjb = "call the bondulance"

func main(){
	x := 69
	// scope of declaration falls below the line of declaration
	fmt.Println(x)
	x = 007
	y := 60 + 9
	fmt.Println(y)
	z := "Bond, James Bond"
	fmt.Println(z)
	fmt.Println(bjb)

	foo ()
}

func foo(){
	fmt.Println("again:",bjb)
}