package file

import (
	"io"
	"os"
	"time"
)

// ISSUE: After the change size might stay the same, use some sort of a hash
func Listen(file *os.File, delay time.Duration) (<-chan []byte, error) {
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	prevSize := info.Size()

	contents := make(chan []byte)

	go func() {
		events := make(chan struct{}, 1)
		for {
			select {
			case <-events:
				content, _ := io.ReadAll(file)
				contents <- content
				file.Seek(0, 0)
			default:
				time.Sleep(delay)
				if info, err := file.Stat(); err == nil && info.Size() != prevSize {
					prevSize = info.Size()
					events <- struct{}{}
				}
			}
		}
	}()

	return contents, nil
}
