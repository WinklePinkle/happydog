package main

import "fmt"

func main() {
	fmt.Println("Hello, everybody! Hi Dr. Nick!")
	foo ()
	fmt.Println("Good news, everyone!")

	for i := 0; i < 100; i++ {
		if i %2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo (){
	fmt.Println("Fight the Foo(ture)")
}
func bar(){
	fmt.Println("Are you still here? It's over. Go home.")
}
// control flow:
// (1) sequence
// (2) loop, iterative
// (3) conditional