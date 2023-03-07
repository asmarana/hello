package main

import (
	"fmt"

	"github.com/asmarana/greetings/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
