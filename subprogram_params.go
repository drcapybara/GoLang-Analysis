package main

import "fmt"

func main() {

	//this works
	printMe(42, "the meaning of life")
	printMe("the meaning of life", 42)

}

func printMe(a_number int, a_string string) {
	fmt.Printf("%v is %s\n", a_number, a_string)
}
