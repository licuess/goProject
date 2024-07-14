package main

import (
	"fmt"
)

func main() {
	foo()
	bar()
	fmt.Println("hello")
	fmt.Println("hello to zed on world")

}

func foo() *int {
	fmt.Println("foo func")
	return nil
}

func bar() {
	fmt.Println("bar func")
	fmt.Println(*foo())
}
