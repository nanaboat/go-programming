package main

import  (
	"fmt"
)

type num int
var x num
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	y = int(x)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
}
