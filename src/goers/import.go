package goers

import "fmt"
import "github.com/jmcvetta/useless"

// Test runs a simple test using the class useless package
// so that we're sure that the goers package is available.
func Test() {
	fmt.Println(useless.Foobar())
}
