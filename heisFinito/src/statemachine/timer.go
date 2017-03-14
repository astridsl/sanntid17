package statemachine

import "time"

var timer_doorOpen time.Time
var timer_orderTimeout time.Time
var timerCounting_doorOpen int = 0
var timerCounting_orderTimeout int = 0 //FINNES DET ET BEDRE NAVN ENN timerCounting?? timerStarted??
var duration_doorOpen float64 = 0
var duration_orderTimeout float64 = 0

func StartTimer_doorOpen() {
	if timerCounting_doorOpen == 0 {
		timer_doorOpen = time.Now() //Setter timeren til det tiden på PC er akkurat nå
		timerCounting_doorOpen = 1  //Betyr at timeren til doorOpen har startet
		duration_doorOpen = 3       //Døren skal stå åpen i 3 sek
	}

}

func StartTimer_orderTimeout(duration float64) {
	timer_orderTimeout = time.Now()
	timerCounting_orderTimeout = 1
	duration_orderTimeout = duration
}

func evIsTimeout(timer int) int {
	//Lager variabler som sjekker om det har gått lengre tid enn det tiumeren er satt til:

	timeDiff_doorOpen := time.Now().Sub(timer_doorOpen).Seconds()
	timeDiff_orderTimeout := time.Now().Sub(timer_orderTimeout).Seconds()

	switch timer {
	case 1:
		if timeDiff_doorOpen >= duration_doorOpen && timerCounting_doorOpen == 1 {
			timerCounting_doorOpen = 0
			return 1
		}

	case 2:
		if timeDiff_orderTimeout >= duration_orderTimeout && timerCounting_orderTimeout == 1 {
			timerCounting_orderTimeout = 0
			return 2
		}
	}
	return 0
}

/*	duration := time.Second * 3
}*/

//var time_t startTime = -1

/*func timer_start(){
	duration := time.Second * 3
	timer := time.NewTimer(duration)

	//startTime := time(0);
}*/

/*func evIsTimeout() int{
	duration := time.Second * 3
	timer := time.NewTimer(duration)

}



/*func doorTimer(timeout chan bool, reset chan bool) {
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
}*/
