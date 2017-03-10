package V1 

import{
	//Packages and directory 
}

func timerToDoor(timeout chan<- bool, reset <-chan bool){ //doorTimer()
	const doorOpenTime = 5*time.Second
	timer := time.NewTimer(0)
	timer.Stop()

	for {
		select {
			
			case <- timer.C: //whats this?? .C?? Block channel C until timer expired  
				timer.Stop()
				timeout <- true 

			case <-reset: 
				timer.Reset(doorOpenTime)
		}
	}
}

func timerToOrder(start <-chan bool, stop <-chan bool, timeout chan<- bool) { //orderTimer()
	const exeOrderTimer = 10*time.Second 
	timer := time.NewTimer(0) 
	timer.Stop() //Se litt mer på timer-funksjoner(funksjonalitet)

	for{
		select {
			case <-start: 
				timer.Reset(exeOrderTimer)

			case <-stop: 
				timer.Stop() 

			case <-timer.C: 
				timer.Stop()
				timeout <- true 
		}
	}
}


