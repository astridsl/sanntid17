package statemachine 

import{
	"./driver"
}

var state config.State
var direction config.MotorDirection
var prevDirr config.MotorDirection
var lastFloor int


func ElevatorManager(){

	//Initialisere driver  

	//finner ut master eller slave? 
	//If master: delegere oppgaver og regne ut kost 
	//If slave: do shit 



}


func initialize(newOrder chan bool, arrivedAtFloor chan int, orderCompleted chan bool, floor int){
	//sette opp channels til timere 
	doorTimeout := make(chan bool)
	doorResetTimer := make(chan, bool)
	orderStartTimer := make(chan, bool)
	orderStopTimer := make(chan, bool)
	orderTimeout := make(chan, bool)

	state = config.Idle
	direction = config.MotorD_stop
	prevDirr = config.MotorD_up
	lastFloor = floor


	/*Queue.Init() 
	
	elev_set_floor_indicator(0)
	currentFloor = elev_get_floor_indicator()

	if(currenFloor == -1){
		Elev_set_motor_direction(config.motorD_down)
		for{
			currentFloor = Elev_get_floor_sensor_signal()
			if currentFloor != -1 {
				break
			}
		}
		Elev_set_motor_direction(config.motorD_stop)
	}

	nextFloor = currentFloor
	lastFloor = currenFloor
	lastFloorStopped = currenFloor
	
	elev_set_floor_indicator(currenFloor)
	motorDirection = Elev_set_motor_direction(config.motorD_stop) 
	state = config.Idle
	currentState = state*/

	go doorTimer(doorTimout, doorResetTimer) 
	go orderTimer(orderStartTimer, orderStopTimer, orderTimeout)  
	go runStateMachine(newOrder, arrivedAtFloor)  //Starter opp statemachine  
}


func runStateMachine(newOrder chan bool, arrivedAtFloor chan int){
	for {
		select {
		case <- newOrder:
			//Funksjon her, noe som gjøre 
			actionNewOrder(orderCompleted, doorTimerReset, doorTimeout, orderStartTimer)

		case f := <- arrivedAtFloor: 
			//Funksjon her 
			actionArrivedAtFloor(doorResetTimer, orderStopTimer, orderCompleted, f)

		case <- doorTimeout: 
			//Funksjon med paramtere her 
			actionDoorTimeout()

		case <- orderTimeout: 
			actionOrderTimeout() 
		}
	}
}

func actionNewOrder( orderCompleted chan int, doorTimerReset chan bool, doorTimeout chan bool, orderStartTimer chan bool) {
	switch state {
	case config.Idle:
		//Går ut ifra at du har fått en ny ordre som du skal utføre.
		direction, prevDirr = queue.ChooseElevDirection(lastFloor, prevDirr)

		if direction == config.MotorD_stop {
			elev.Elev_set_door_open_lamp(1)
			doorTimerReset <- true 
			queue.DelLocalOrderAtFloor(lastFloor)
			orderCompleted <- true
			state = config.DoorOpen 

		} else {
			elev.Elev_set_motor_direction(direction)//Heisen kjører!!!
			orderStartTimer <- true
			state = config.Moving


		}

	case config.DoorOpen:
		direction, prevDirr = queue.ChooseElevDirection(lastFloor, prevDirr)
		if direction == config.MotorD_stop {
			doorTimerReset <- true
			queue.DelLocalOrderAtFloor(lastFloor)
			orderCompleted <- true
		}

	default:
			//Don't do anything 
	} 

}

func actionArrivedAtFloor(doorResetTimer chan bool, orderStopTimer chan bool, orderCompleted chan bool, newFloor int){
	lastFloor = newFloor
	switch state {
	case config.Moving: 
		elev.Elev_set_motor_direction(config.MotorD_stop)
		elev.Elev_set_door_open_lamp(1) 
		doorResetTimer <- true 
		orderStopTimer <- true
		orderCompleted <- true

		state = config.Idle 

	case config.DoorOpen: 
		elev.Elev_set_door_open_lamp(0)
		orderCompleted <- true
		state = config.Idle 

	default: 
		//Don't do anything 
	}
}

func actionDoorTimeout(){
	elev.Elev_set_door_open_lamp(0)
	state = config.Idle 

}


func actionOrderTimeout(){
	elev.Set_motor_direction(MotorD_stop)
	elev.Set_door_open_light(0)

	queue.DelLocalOrderAtFloor(lastFloor)
	//Skal vi ha med dette?
	for button := 0; button < config.N_BUTTONS; button++ {
		elev.button_channel_matrix[lastFloor][button]
	}

}








