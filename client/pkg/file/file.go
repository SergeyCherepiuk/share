package file

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"sync"
	"time"

	"github.com/SergeyCherepiuk/share/client/pkg/clean"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/med"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/ot"
)

// File structure manages content of an OS file
// and exposes two chanels for receiving and outputing changes
type File struct {
	In  chan ot.Operation
	Out chan ot.Operation

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
		In:      make(chan ot.Operation),
		Out:     make(chan ot.Operation),
		path:    path,
		content: make([]byte, 0),
	}

	go file.watch(100 * time.Millisecond)
	go file.apply()

	return &file, nil
}

// Watches the OS file for changes, computes an OT-ot
// and sends the operations to file's out channel
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

			for _, operation := range ot.Adjust(med.Diff(prev, f.content)) {
				f.Out <- operation
			}
		}
		time.Sleep(delay)
	}
}

// Accepts the operations from file's in channel,
// applies them to file's underlying slice (content)
// and tries to write an updated content to an OS file
func (f *File) apply() {
	for {
		operation := <-f.In

		// TODO: Handle errors
		switch operation.Type {
		case ot.INSERTION:
			f.insert(operation.Character, operation.Position)
		case ot.DELETION:
			f.delete(operation.Position)
		case ot.SUBSTITUTION:
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
	var err error
	errors.Join(err, os.WriteFile(f.path, f.content, 0644))
	return err
}
