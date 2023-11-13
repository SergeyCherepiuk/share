package ot

import (
	"fmt"
)

// var Operations = make(chan diff.Operation)
var Operations = make(chan string)

func Apply(path string) {
	for operation := range Operations {
		fmt.Printf("Appling %+v to %s\n", operation, path)
	}
}
