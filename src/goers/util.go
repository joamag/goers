package goers

import "code.google.com/p/go-uuid/uuid"

func UUID() (string, error) {
	return uuid.New(), nil
}
