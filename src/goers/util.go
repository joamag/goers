//+build full

package goers

import "github.com/pborman/uuid"

// UUID generates a UUID value using an external library
// creating a simple yet useful abstraction for such operation.
func UUID() (string, error) {
    return uuid.New(), nil
}
