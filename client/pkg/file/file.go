package file

import (
	"os"
	"time"
)

func Listen(path string, delay time.Duration) (<-chan []byte, error) {
	info, err := os.Stat(path)
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
				content, _ := os.ReadFile(path) // TODO: Handle the error
				contents <- content
			default:
				if info, err := os.Stat(path); err == nil && !info.ModTime().Equal(prevModTime) {
					prevModTime = info.ModTime()
					events <- struct{}{}
				}
				time.Sleep(delay)
			}
		}
	}()

	return contents, nil
}
