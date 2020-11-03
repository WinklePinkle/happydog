package main

import "fmt"

type batman int

var x batman
var y int

func main(){
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 187
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)

}
