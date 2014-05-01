package goers

import "os"
import "fmt"

func Execute() {
	args := os.Args;
	host := DEFAULT_HOST
	if len(args) > 1 { host = args[1] }
	err := Get(host)
	if err != nil { fmt.Println(err) }
}

func main() {
	Execute()
}
