/*
	go mod edit -replace="github.com/greetings"="../greetings"
	go mod tidy
*/

package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Cshouu")
	fmt.Println(message)
}
