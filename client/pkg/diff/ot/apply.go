package ot

import (
	"fmt"

	"github.com/SergeyCherepiuk/share/client/pkg/diff"
)

var Operations = make(chan diff.Operation)

func Apply(path string) {
	for operation := range Operations {
		fmt.Printf("Appling %+v to %s\n", operation, path)
	}
}
