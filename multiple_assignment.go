package main

import "fmt"

func multiple_return(x int, y int) (int, int) {
	return x + y, y * x
}

func main() {
	first, second := 5, 5
	third, fourth := multiple_return(first, second)
	fmt.Printf("third: %v \nfourth: %v\n", third, fourth)
}
