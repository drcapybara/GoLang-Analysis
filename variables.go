package main

//instantiate a new type
type exampleType struct {
	example string
}

//modify a field from the type passed in
func (t exampleType) exampleFunc() {
	t.example = "Am I different?"
}

func main() {
	var exampleTest = exampleType{example: "Here I am, here I remain"}
	exampleTest.exampleFunc()
	println(exampleTest.example)
}
