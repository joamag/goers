package main

import "net"
import "fmt"
import "bufio"

func get(host string) {
	// prints a debug information into the current
	// command line output indicating the retrieval
	// of information from the remote server
	fmt.Printf("GET http://%s\n", host)

	// establishes the connection to the remote host,
	// indicating the problem in case it exists
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println("There was an error with connection")
		return
	}

	// tries to run a simple get statement on the
	// current connection, this is hardcoded to
	// retrieve the root element of the current element
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	buffer := bufio.NewReader(conn)
	status, err := buffer.ReadString((byte)('\n'))
	if err != nil {
		fmt.Println("There was an error in status reading")
		return
	}

	// prints the initial status line of the received
	// request to the standard output buffer
	fmt.Printf(status)

	// iterates continuously printing the various bytes
	// that are currently being received
	for true {
		// reads a line from the current buffer and in
		// case the length of the received bytes is
		// zero, then end od communication has been
		// reached and must break the current loop
		line, _, err := buffer.ReadLine()
		if len(line) == 0 {
			break
		}

		// in case the error flag is currently set, must print
		// a message indicating the problem that was just "raised"
		if err != nil {
			fmt.Println("There was an error in line reading")
			return
		}

		// creates a string value out of the line
		// and then prints to the standard output
		lineS := string(line[:])
		fmt.Printf(lineS + "\n")
	}
}

func main() {
	get("google.com:80")
}
