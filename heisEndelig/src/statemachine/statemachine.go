package statemachine

import (
	"../config"
	"../driver"
	"../queue"
	"fmt"
	"time"
)

var state config.State
var direction config.MotorDirection
var prevDirr config.MotorDirection
var lastFloor int

/*
func ElevatorManager(){

	//Initialisere driver

	//finner ut master eller slave?
	//If master: delegere oppgaver og regne ut kost
	//If slave: do shit



}
*/

func Initialize(newOrder chan bool, arrivedAtFloor chan bool, floor int) {
	//sette opp channels til timere
	doorTimeout := make(chan bool)
	doorResetTimer := make(chan bool)
	orderStartTimer := make(chan bool)
	orderStopTimer := make(chan bool)
	orderTimeout := make(chan bool)

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

	fmt.Println("state is:", state)

	go updateLastFloor() //Oppdaterer hele tiden den siste etasjen heisen befant seg i. DENNE FUNGERER!
	go doorTimer(doorTimeout, doorResetTimer)
	go orderTimer(orderStartTimer, orderStopTimer, orderTimeout)
	go runStateMachine(newOrder, arrivedAtFloor, doorResetTimer, doorTimeout, orderStartTimer, orderStopTimer, orderTimeout) //Starter opp statemachine
}

func runStateMachine(newOrder chan bool, arrivedAtFloor chan bool, doorResetTimer chan bool, doorTimeout chan bool, orderStartTimer chan bool, orderStopTimer chan bool, orderTimeout chan bool) {
	for {
		select {
		case <-newOrder:
			//Funksjon her, noe som gjøre
			fmt.Println("The elevator has gotten a new order to take")
			actionNewOrder(doorResetTimer, doorTimeout, orderStartTimer, arrivedAtFloor, orderTimeout)

		case <-arrivedAtFloor:
			//Funksjon her
			actionArrivedAtFloor(doorResetTimer, orderStopTimer, newOrder)

		case <-doorTimeout:
			//Funksjon med paramtere her
			//actionDoorTimeout(timeOut, reset)

		case <-orderTimeout:
			//actionOrderTimeout(start, stop, timeOut)
		}
		time.Sleep(15 * time.Millisecond)
	}
}

func actionNewOrder(doorTimerReset chan bool, doorTimeout chan bool, orderStartTimer chan bool, arrivedAtFloor chan bool, orderTimeout chan bool) {
	switch state {
	case config.Idle:
		//Går ut ifra at du har fått en ny ordre som du skal utføre.
		direction, prevDirr = queue.ChooseElevDirection(lastFloor, prevDirr)

		if direction == config.MotorD_stop || queue.ElevShouldStop(lastFloor, direction) {
			fmt.Println("elev should stop")
			driver.Elev_set_door_open_lamp(1)
			driver.Elev_set_motor_direction(config.MotorD_stop)
			doorTimerReset <- true
			queue.DelLocalOrderAtFloor(lastFloor)
			state = config.DoorOpen
		} else {
			fmt.Println("elev should drive")

			driver.Elev_set_motor_direction(direction) //Heisen kjører!!!
			fmt.Println("?????????")
			/*if queue.IsOrderAtFloor(lastFloor) {
				driver.Elev_set_motor_direction(config.MotorD_stop)
				arrivedAtFloor <- true
			}*/
			//orderTimeout <- true
			//orderStartTimer <- true
			//queue.ElevShouldStop(lastFloor, direction)
			arrivedAtFloor <- true
			state = config.Moving
		}

	case config.DoorOpen: //copy and paste absolute sign
		direction, prevDirr = queue.ChooseElevDirection(lastFloor, prevDirr)
		if direction == config.MotorD_stop {
			doorTimerReset <- true
			queue.DelLocalOrderAtFloor(lastFloor)
		}

	case config.Moving:
		fmt.Println("Inni moving!!!!")
		direction, prevDirr = queue.ChooseElevDirection(lastFloor, prevDirr)
		if queue.ElevShouldStop(lastFloor, direction) {
			driver.Elev_set_motor_direction(config.MotorD_stop)
			//arrivedAtFloor <- lastFloor
		}

	default:
		//Don't do anything
		fmt.Println("default")
	}
}

func updateLastFloor() {
	for {
		if driver.Elev_get_floor_sensor_signal() != -1 {
			lastFloor = driver.Elev_get_floor_sensor_signal()
			driver.Elev_set_floor_indicator(lastFloor) //Sett lys på etasjen heisen er i.

		}
		time.Sleep(15 * time.Millisecond)
	}
}

func actionArrivedAtFloor(doorResetTimer chan bool, orderStopTimer chan bool, newOrder chan bool) {
	//lastFloor = newFloor
	switch state {
	case config.Moving:
		fmt.Println("er i floor arrived")
		driver.Elev_set_motor_direction(config.MotorD_stop)
		driver.Elev_set_door_open_lamp(1)
		queue.DelLocalOrderAtFloor(lastFloor)
		doorResetTimer <- true
		orderStopTimer <- true

		state = config.Idle

	case config.DoorOpen:
		driver.Elev_set_door_open_lamp(0)
		state = config.Idle

	default:
		//Don't do anything
	}

	for floor := 0; floor < config.N_FLOORS; floor++ {
		for button := 0; button < config.N_BUTTONS; button++ {
			if queue.LocalQueue[floor][button] == 1 {
				newOrder <- true
			} else {
				println("No orders in queue")
			}
		}
	}
}

func actionDoorTimeout() {
	driver.Elev_set_door_open_lamp(0)
	state = config.Idle

}

func actionOrderTimeout() {
	driver.Elev_set_motor_direction(config.MotorD_stop)
	driver.Elev_set_door_open_lamp(0)

	queue.DelLocalOrderAtFloor(lastFloor)

}
