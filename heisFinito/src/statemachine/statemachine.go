package statemachine

import (
	"../config"
	"../driver"
	"../queue"
	"fmt"
	"time"
)

var MotorDirection int

//var MotorDirection config.MotorDirection
var NextFloor int
var LastFloor int
var LastFloorStopped int

var CurrentFloor int
var ButtonPressed int
var AtSelectedFloor int
var TimeOut int
var StopSignalPressed int
var CurrentState config.State

func Initialize() {

	var init_success int = driver.Io_init()

	if init_success == 0 {
		fmt.Println("Initialization failed")
	}

	queue.DelQueue()
	driver.Elev_set_floor_indicator(0)
	CurrentFloor = driver.Elev_get_floor_sensor_signal()
	fmt.Println("Current floor is", CurrentFloor)
	driver.Elev_set_motor_direction(-1)

	/*if CurrentFloor == -1 {
		driver.Elev_set_motor_direction(-1)
		for CurrentFloor == -1 {
			CurrentFloor = driver.Elev_get_floor_sensor_signal()
		}
		driver.Elev_set_motor_direction(0)
	}*/

	if CurrentFloor == -1 {
		driver.Elev_set_motor_direction(-1)
		for {
			CurrentFloor = driver.Elev_get_floor_sensor_signal()
			//fmt.Println("Looping")
			if CurrentFloor != -1 {
				break
			}
		}
		driver.Elev_set_motor_direction(0)
	}

	NextFloor = CurrentFloor
	LastFloor = CurrentFloor
	LastFloorStopped = CurrentFloor
	driver.Elev_set_floor_indicator(LastFloor)
	MotorDirection = -1
	CurrentState = config.State_idle
}

func evButtonSignalPressed() int {
	for floor := 0; floor < config.N_FLOORS; floor++ {
		if floor != 3 && driver.Elev_get_button_signal(config.Button_Up, floor) == 1 {
			queue.AddOrderToLocalQueue(floor, 0)
			return 1
		} else if floor != 0 && driver.Elev_get_button_signal(config.Button_Down, floor) == 1 {
			queue.AddOrderToLocalQueue(floor, 1)
			return 1
		} else if driver.Elev_get_button_signal(config.Button_Command, floor) == 1 {
			queue.AddOrderToLocalQueue(floor, 2)
			return 1
		}
	}
	return 0
}

func evAtSelectedFloor() int {
	for button := 0; button < config.N_BUTTONS; button++ {
		if driver.Elev_get_floor_sensor_signal() == NextFloor && queue.GetButtonPressed(NextFloor, button) == 1 {
			return 1
		}
	}
	if driver.Elev_get_floor_sensor_signal() == NextFloor {
		return -1
	}
	return 0
}

func ChooseMotorDirection() int {
	if LastFloor == NextFloor || NextFloor == -1 {
		driver.Elev_set_motor_direction(0)
		if LastFloor == config.N_FLOORS-1 {
			return -1
		} else if LastFloor == 0 {
			return 1
		}
		return MotorDirection
	} else if LastFloor < NextFloor && CurrentState != config.State_door_open {
		driver.Elev_set_motor_direction(1)
		return 1
	} else if LastFloor > NextFloor && CurrentState != config.State_door_open {
		driver.Elev_set_motor_direction(-1)
		return -1
	}
	return MotorDirection
}

func EventManager() {
	StopSignalPressed = driver.Elev_get_stop_signal()
	if StopSignalPressed == 0 {
		ButtonPressed = evButtonSignalPressed()
		AtSelectedFloor = evAtSelectedFloor()
		CurrentFloor = driver.Elev_get_floor_sensor_signal()
		MotorDirection = ChooseMotorDirection()
		NextFloor = queue.GetNextFloor(LastFloorStopped, MotorDirection, LastFloor, NextFloor)
		//Timeout = evIsTimeout(); //SJEKK DENNE!
		if CurrentFloor != -1 {
			LastFloor = CurrentFloor
			driver.Elev_set_floor_indicator(LastFloor)
		}
		//Følgende kan bare gjelde dersom heisen har vært stoppet. Sjekker om neste etasje er den forige etasjen heisen kjørte forbi slik at man skal kunne kjøre tilbake til denne etasjen dersom den bestilles.
		if NextFloor == LastFloor && CurrentFloor == -1 {
			if MotorDirection == 1 {
				LastFloor++
			} else {
				LastFloor--
			}
		}
	} else {
		driver.Elev_set_stop_lamp(1)
		driver.Elev_set_motor_direction(0)
		queue.DelQueue()
		CurrentState = config.State_emergency_stop
		NextFloor = -1 //Indikerer at heisen skal stå stille helt til den får en ny bestilling.
	}
}

func DriveElevator() {
	switch CurrentState {
	case config.State_idle:
		if AtSelectedFloor == 1 {
			//timer_start() SE PÅ DENNE!
			driver.Elev_set_door_open_lamp(1)
			queue.DelAllOrdersInFloor(LastFloor)
			LastFloorStopped = LastFloor
			CurrentState = config.State_door_open
		} else if AtSelectedFloor == 0 {
			CurrentState = config.State_moving
		}
	case config.State_moving:
		if AtSelectedFloor == 1 {
			//timer_start();
			driver.Elev_set_door_open_lamp(1)
			queue.DelAllOrdersInFloor(LastFloor)
			LastFloorStopped = LastFloor
			CurrentState = config.State_door_open
		}
	case config.State_door_open:
		/*if (timeout == 1){
			elev_set_door_open_lamp(0);
			current_state = state_idle;
		}else if (buttonPressed == 1 && nextFloor == lastFloorStoped){
			timer_start();
			queue_del_one_floor();
		}*/
		//LITT JALLA!!
		time.Sleep(time.Second * 3)
		driver.Elev_set_door_open_lamp(0)
		CurrentState = config.State_idle

	case config.State_emergency_stop:
		if StopSignalPressed == 0 && CurrentFloor != -1 {
			driver.Elev_set_stop_lamp(0)
			driver.Elev_set_door_open_lamp(0)
			CurrentState = config.State_idle
		} else if CurrentFloor != -1 {
			driver.Elev_set_door_open_lamp(1)
		} else if StopSignalPressed == 0 {
			driver.Elev_set_stop_lamp(0)
			CurrentState = config.State_undefined
		}
	case config.State_undefined:
		if ButtonPressed == 1 {
			CurrentState = config.State_moving
		}
	}
}
