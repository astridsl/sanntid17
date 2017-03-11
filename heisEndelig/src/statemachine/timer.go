package statemachine 

import{
	"time"
	//Directory 
}


 func doorTimer(timeout chan bool, reset chan bool){
 	duration = time.Second *3
 	timer = time.NewTimer(duration)

 	select {
 		case <- reset: 
 			time.Reset(duration)

		case <- time.C: 
			timeout <- true   
 	} 

 	time.Stop()
 }

 func orderTimer(start chan bool, stop chan bool, timeout chan bool){
	duration = time.Second *3
 	timer = time.NewTimer(duration)

 	select {
 		case <- start: 
 			time.Reset(duration)

		case <- stop: 
			time.stop()

		case <- time.C: 
			timeout <- true   
 	} 

 	time.Stop()
 }