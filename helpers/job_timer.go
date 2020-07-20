package helpers

import (
	"time"
)

// DoJobEvery executes given func every given duration
func DoJobEvery(d time.Duration, job func()) {
	ticker := time.NewTicker(d * time.Second)
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				ticker.Stop()
				return
			case <-ticker.C:
				job()
			}
		}
	}()

}
