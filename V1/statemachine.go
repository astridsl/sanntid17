package V1 

//NB: MÅ HUSKE Å ENDRE DE VARIABLENE SOM HØRER TIL config (HAR def. FORANN). KAN VÆRE AT VI MÅ TA BORT def.

import{
	
	//Directories 
}

//Defining variables (using variables from config (driver))
var state def.State
var dir def.MotorDirection
var prevDir def.MotorDirection
var lastFloor int

func initialize() {
	//Make channels for door and order timer functions
	//Order timer:
	startOrderTimer := make(chan, bool)
	endOrderTimer := make(chan, bool)
	orderTimeout := make(chan, bool)
	//Door timer:
	doorTimeout := make(chan, bool)
	doorResetTimer := make(chan, bool)

	//Choose which "positions" the variables should be in
	state = def.Idle
	dir = def.MD_stop
	prevDir = def.MD_down
	lastFloor = floor



	//Run go routines. (Handle door timer and order timer, and run FSM)
	go orderTimer(startOrderTimer, endOrderTimer, orderTimeout)
	go doorTimer(doorTimeout, doorResetTimer)
	go runStateMachine()



}

func runStateMachine(arrivedAtFloor, orderCompleted chan int, newOrder, startOrderTimer, endOrderTimer, orderTimeout chan bool, doorTimeout, doorResetTimer, statusUpdate){
	for {
		select {
			//Handle new orders
			case <-newOrder:
				actionNewOrder()

			//floor reached --> order completed
			case <-arrivedAtFloor:
				actionArrivedAtFloor()

			//order timeout --> it takes to loong time before the orer is handled. --> stophandleing the order and give it to another elevator
			case <-orderTimeout:
				actionOrderTimeout()

			//The door has been open too long --> Close it and restart timer?
			case <-doorTimeout:
				actionDoorTimeout()
		}
	}


}

func actionNewOrder(orderCompleted chan int, startOrderTimer chan bool, doorResetTimer){
	switch state {
	case def.Idle:

	case def.DoorOpen:

	default:
		//The elevator is already handeling an order. --> Ignore.
	}


}

func actionArrivedAtFloor(){


}

func actionDoorTimeout(){

}

func actionOrderTimeout() {
	
}

