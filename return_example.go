package main

import "fmt"

func returnToSender(delivery string) (content string) {
	content = "bird" + delivery
	return
}

func main() {
	payload := returnToSender(" golf")
	fmt.Println(payload)
}
