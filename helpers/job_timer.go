package helpers

import "time"

// DoJobEvery executes given func every given duration
func DoJobEvery(d time.Duration, job func()) {
	for range time.Tick(d) {
		job()
	}
}
