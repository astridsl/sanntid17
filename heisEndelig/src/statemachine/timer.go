package statemachine

import (
	"time"
)

func doorTimer(timeout chan bool, reset chan bool) {
	duration := time.Second * 3
	timer := time.NewTimer(duration)

	select {
	case <-reset:
		timer.Reset(duration)

	case <-timer.C:
		timeout <- true
	}

	timer.Stop()
}

func orderTimer(start chan bool, stop chan bool, timeout chan bool) {
	duration := time.Second * 3
	timer := time.NewTimer(duration)

	select {
	case <-start:
		timer.Reset(duration)

	case <-stop:
		timer.Stop()

	case <-timer.C:
		timeout <- true
	}

	timer.Stop()
}
