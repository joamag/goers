package main

import (
	"fmt"

	"github.com/joamag/goers/src/goers"
)

func main() {
	var message = "Running goers..."
	fmt.Println(message)
	goers.Execute()
}
