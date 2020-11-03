package main

import (
	"fmt"
)

var x = "gonna give it to ya"
var a = "ex gonna give it to ya"
var y = 007

func main (){
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 69
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x",y, y, y)
}