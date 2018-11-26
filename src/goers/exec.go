package goers

import "os"
import "fmt"

// Execute runs the default execution process that tries to
// retrieve HTTP information based on the command line
// arguments and fallback to the static values otherwise.
func Execute() {
    args := os.Args
    host := DefaultHost
    if len(args) > 1 {
        host = args[1]
    }
    err := Get(host)
    if err != nil {
        fmt.Println(err)
    }
}
