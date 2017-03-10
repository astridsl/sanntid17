package V1 

//NB: MÅ HUSKE Å ENDRE DE VARIABLENE SOM HØRER TIL config (HAR def. FORANN). KAN VÆRE AT VI MÅ TA BORT def.

import{
	
	//Directories 
}

//Defining variables (using variables from config (driver))
var state def.State
var direction def.MotorDirection
var prevDirection def.MotorDirection
var lastFloor int

func initialize(arrivedAtFloor, orderCompleted chan int, newOrder, doorTimeout, doorTimerReset, updateStatus, startOrderTimer, endOrderTimer, orderTimeout chan bool){
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
	direction = def.MotorD_stop
	prevDirection = def.MotorD_down
	lastFloor = floor



	//Run go routines. (Handle door timer and order timer, and run FSM)
	go orderTimer(startOrderTimer, endOrderTimer, orderTimeout)
	go doorTimer(doorTimeout, doorResetTimer)
	go runStateMachine(arrivedAtFloor, orderCompleted chan int, newOrder, startOrderTimer, endOrderTimer, orderTimeout chan bool, doorTimeout, doorResetTimer, statusUpdate)



}

func runStateMachine(arrivedAtFloor, orderCompleted chan int, newOrder, startOrderTimer, endOrderTimer, orderTimeout chan bool, doorTimeout, doorResetTimer, statusUpdate){
	for {
		select {
			//Handle new orders
			case <-newOrder:
				actionNewOrder(orderCompleted, doorTimerReset, startOrderTimer)

			//floor reached --> order completed
			case f := <-arrivedAtFloor:
				actionArrivedAtFloor(f, orderCompleted, doorTimerReset, stopOrderTimer)

			//order timeout --> it takes to loong time before the orer is handled. --> stophandleing the order and give it to another elevator
			case <-orderTimeout:
				//actionOrderTimeout() SKRIV INN KODEN I DENNE FUNKSJONEN
				//MÅ ENDRE PÅ DENNE KODEN!!!!!
				driver.SetMotorDirection(def.MD_stop)
				for f := 0; f < def.N_FLOORS; f++ {
					for b := 0; b < def.N_BUTTONS; b++ {
						driver.SetButtonLight(f, b, false)
					}
				}
				driver.SetDoorOpenLight(false)
				fmt.Printf("Order has timed out. Terminate program. ")
				os.Exit(1)

			//The door has been open too long --> Close it and restart timer?
			case <-doorTimeout:
				actionDoorTimeout(startOrderTimer)
		}
		updateStatus <- true
	}


}

func actionNewOrder(orderCompleted chan int, startOrderTimer chan bool, doorResetTimer){
	switch state {
	case def.Idle:
		direction, prevDirection = Queue.ChooseElevDirection(lastFloor, prevDirection)
		if Queue.ElevShouldStop(lastFloor, direction) {
			doorResetTimer  <- true
			Queue.DelLocalOrdersAtFloorFromQueue(lastFloor)
			orderCompleted <- lastFloor
			elev.SetDoorOpenLight(true)
			state=def.DoorOpen
		}
		else {
			elev.SetMotorDirection(direction)
			state = def.Moving
			startOrderTimer <- true
		}


	case def.DoorOpen:
		if  Queue.ElevShouldStop(lastFloor, direction){
			startOrderTimer <- true
			Queue.DelLocalOrdersAtFloorFromQueue(lastFloor)
			orderCompleted <- lastFloor
		}

	default:
		//The elevator is already handeling an order. --> Ignore.
	}


}

func actionArrivedAtFloor(newFloor int, orderCompleted chan int, doorTimerReset, stopOrderTimer chan bool){
lastFloor = newFloor
elev.SetFloorIndicator(lastFloor)
switch state {
case def.Moving:
	if Queue.ElevShouldStop(lastFloor, direction) {
		orderCompleted <- lastFloor
		Queue.DelLocalOrdersAtFloorFromQueue(lastFloor)
		direction=def.MotorD_stop
		elev.SetMotorDirection(dir)
		elev.SetDoorOpenLight(true)
		state = def.DoorOpen
		doorTimerReset <- true
	}
default:

}

}

//DENNE FUNKSJONENE MÅ ENDRES!!
func actionDoorTimeout(startOrderTimer chan bool){
	switch state {
	case def.DoorOpen:
		driver.SetDoorOpenLight(false)
		dir, prevDir = queue.ChooseDirection(lastFloor, prevDir)
		driver.SetMotorDirection(dir)
		if dir == def.MD_stop {
			state = def.Idle
		} else {
			state = def.Moving
			startOrderTimer <- true
		}

	default:
		// Ignore
	}
}

func actionOrderTimeout() {
	
}
