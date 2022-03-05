package main

import "fmt"

func main() {

	type apple struct{ appleType string }
	type orange struct{ orangeType string }

	type experimentalFruit struct {
		apple
		orange
	}
	fruit_basket := experimentalFruit{
		apple{appleType: "cosmic crisp"},
		orange{orangeType: "tangelo"},
	}
	fmt.Println(fruit_basket.apple)
	fmt.Println(fruit_basket.orange)
}
