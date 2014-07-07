package goers

import "code.google.com/p/go-uuid/uuid"

// Generates a UUID value using an external library creating
// a simple yet usefull abstraction for such operation.
func UUID() (string, error) {
    return uuid.New(), nil
}
