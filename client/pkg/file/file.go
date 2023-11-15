package file

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"sync"
	"time"

	"github.com/SergeyCherepiuk/share/client/pkg/clean"
	"github.com/SergeyCherepiuk/share/client/pkg/diff"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/ot"
)

const format = "2006-01-02 15:04:05.000000000 -0700"

// File structure manages content of an OS file
// and exposes two chanels for receiving and outputing changes
type File struct {
	In  chan diff.Operation
	Out chan diff.Operation

	path string

	muContent sync.RWMutex
	content   []byte
}

func New(path string, preserve bool) (*File, error) {
	if _, err := os.Stat(path); err == nil {
		return nil, fmt.Errorf("file already exists")
	}

	if _, err := os.Create(path); err != nil {
		return nil, err
	}
	if !preserve {
		clean.Add(func() { os.Remove(path) })
	}

	file := File{
		In:      make(chan diff.Operation),
		Out:     make(chan diff.Operation),
		path:    path,
		content: make([]byte, 0),
	}

	go file.watch(100 * time.Millisecond)
	go file.apply()

	return &file, nil
}

func (f *File) watch(delay time.Duration) {
	info, _ := os.Stat(f.path)
	prevModTime := info.ModTime()

	for {
		if info, err := os.Stat(f.path); err == nil && !info.ModTime().Equal(prevModTime) {
			prev := f.content
			prevModTime = info.ModTime()

			f.muContent.Lock()
			f.content, _ = os.ReadFile(f.path)
			f.muContent.Unlock()

			for _, operation := range ot.Diff(prev, f.content) {
				f.Out <- operation
			}
		}
		time.Sleep(delay)
	}
}

// TODO: Handle errors
func (f *File) apply() {
	for {
		operation := <-f.In
		switch operation.Type {
		case diff.INSERTION:
			f.insert(operation.Character, operation.Position)
		case diff.DELETION:
			f.delete(operation.Position)
		case diff.SUBSTITUTION:
			f.substitute(operation.Character, operation.Position)
		}
	}
}

func (f *File) insert(b byte, at int) error {
	f.muContent.Lock()
	f.content = slices.Insert(f.content, at, b)
	f.muContent.Unlock()
	return f.save()
}

func (f *File) delete(at int) error {
	f.muContent.Lock()
	f.content = slices.Delete(f.content, at, at+1)
	f.muContent.Unlock()
	return f.save()
}

func (f *File) substitute(b byte, at int) error {
	f.muContent.Lock()
	f.content[at] = b
	f.muContent.Unlock()
	return f.save()
}

func (f *File) save() error {
	info, _ := os.Stat(f.path)
	touch := exec.Command("touch", "-m", "-d", info.ModTime().Format(format), f.path)

	var err error
	errors.Join(err, os.WriteFile(f.path, f.content, os.ModeAppend))
	errors.Join(err, touch.Run())
	return err
}
