package file

import (
	"io"
	"os"
	"time"
)

func Listen(file *os.File, delay time.Duration) (<-chan []byte, error) {
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	prevModTime := info.ModTime()

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
				if info, err := file.Stat(); err == nil && !info.ModTime().Equal(prevModTime) {
					prevModTime = info.ModTime()
					events <- struct{}{}
				}
			}
		}
	}()

	return contents, nil
}
