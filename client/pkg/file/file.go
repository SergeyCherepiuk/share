package file

import (
	"io"
	"os/exec"

	"github.com/SergeyCherepiuk/share/client/pkg/clean"
)

// ISSUEs:
//  1. For huge files sends content in several separate messages
//  2. Isn't triggered if file is emptied
//  3. Hard to read (refactor needed)
func Listen(path string) (<-chan []byte, error) {
	cmd := exec.Command("tail", "-f", path)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	clean.Add(func() { stdout.Close() })

	fileParts := make(chan []byte, 100)
	go read(stdout, fileParts)

	fileContents := make(chan []byte)

	go func() {
		content := make([]byte, 0)
		for {
			select {
			case part := <-fileParts:
				content = append(content, part...)
			default:
				if len(content) != 0 {
					fileContents <- content
					content = make([]byte, 0)
				}
			}
		}
	}()

	go cmd.Run()
	return fileContents, nil
}

func read(reader io.Reader, ch chan<- []byte) {
	buf := make([]byte, 512)
	for {
		n, _ := reader.Read(buf)
		ch <- buf[:n]
	}
}
