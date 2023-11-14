package ot

import (
	"os"
	"os/exec"

	"github.com/SergeyCherepiuk/share/client/pkg/diff"
)

const format = "2006-01-02 15:04:05.000000000 -0700"

var Operations = make(chan diff.Operation)

func Apply(path string) {
	info, _ := os.Stat(path) // TODO: Handle as error
	mt := info.ModTime()

	for operation := range Operations {
		operation.Apply(path) // TODO: Handle an error
		exec.Command("touch", "-m", "-d", mt.Format(format), path).Run()
	}
}
