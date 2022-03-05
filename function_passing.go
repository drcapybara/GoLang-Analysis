package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

/*contrary to other familiar languages, the
function signature comes after the argument list*/
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}
func main() {
	partial := partialSum(3)
	fmt.Println(partial(7))
}
