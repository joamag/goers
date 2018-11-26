package goers

import "net"
import "fmt"
import "bufio"
import "strings"

// DefaultPort defines the port that is going to be used in
// case no other HTTP port is defined
const DefaultPort = 80

// DefaultHost defines the host that is going to be used in
// case no other host is provided
const DefaultHost = "google.com:80"

// Get retrieves the header contents from an HTTP based connection
// with the provided host string, the string should conform with
// the HOST:PORT format to be able to be used correctly.
func Get(host string) error {
    // splits the provided host string into the host an port
    // parts and in case there's no port part it's added using
    // the default port value defined in the package
    parts := strings.Split(host, ":")
    if len(parts) < 2 {
        host = fmt.Sprintf("%s:%d", host, DefaultPort)
    }

    // prints a debug information into the current
    // command line output indicating the retrieval
    // of information from the remote server
    fmt.Printf("GET http://%s\n", host)

    // establishes the connection to the remote host,
    // indicating the problem in case it exists
    conn, err := net.Dial("tcp", host)
    if err != nil {
        return err
    }

    // tries to run a simple get statement on the
    // current connection, this is hardcoded to
    // retrieve the root element of the current element
    _, err = fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    if err != nil {
        return err
    }
    buffer := bufio.NewReader(conn)
    status, err := buffer.ReadString((byte)('\n'))
    if err != nil {
        return err
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
        if err != nil {
            return err
        }
        if len(line) == 0 {
            break
        }

        // creates a string value out of the line
        // and then prints to the standard output
        lineS := string(line[:])
        fmt.Printf(lineS + "\n")
    }

    // returns an empty error value meaning that no error
    // as occurred for the current session
    return nil
}
