package main

import (
	"fmt"
	"time"
)

func thread_test(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	thread_test("direct")

	go thread_test("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("go-ing lol")

	time.Sleep(time.Second)
	fmt.Println("done")
}
