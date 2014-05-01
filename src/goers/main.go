package goers

import "os"
import "fmt"

func main() {
	args := os.Args;
	host := DEFAULT_HOST
	if len(args) > 1 { host = args[1] }
	err := get(host)
	if err != nil { fmt.Println(err) }
}
