package driver

import (
	"../config"
	//"fmt"
	//"time"
)

const MOTOR_SPEED = 2800

//const N_FLOORS = 4
//const N_BUTTONS = 3

var lamp_channel_matrix = [config.N_FLOORS][config.N_BUTTONS]int{
	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var Button_channel_matrix = [config.N_FLOORS][config.N_BUTTONS]int{
	{BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
	{BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
	{BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
	{BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

/*type ElevButtonType int

const (
	BUTTON_CALL_UP ElevButtonType = iota
	BUTTON_CALL_DOWN
	BUTTON_COMMAND
)*/

/*func Elev_init() int {

	var init_success int = Io_init()
	for f := 0; f < config.N_FLOORS; f++ {
		for b := 0; b < config.N_BUTTONS; b++ {
			Elev_set_button_lamp(b, f, 0)
		}
	}

	fmt.Println("Kommer vi hit?")
	Elev_set_stop_lamp(0)
	Elev_set_door_open_lamp(0)
	//Elev_set_floor_indicator(Elev_get_floor_sensor_signal())

	//Elev_set_floor_indicator(0)

	fmt.Println("Kommer vi hit da??")
	currentFloor := Elev_get_floor_sensor_signal()

	if currentFloor == -1 {
		Elev_set_motor_direction(-1)
		for {
			currentFloor = Elev_get_floor_sensor_signal()
			if currentFloor != -1 {
				break
			}
		}
		Elev_set_motor_direction(0)
	}

	Elev_set_floor_indicator(currentFloor)
	//motorDirection = Elev_set_motor_direction(config.motorD_stop)
	//state = config.Idle

	return init_success
}*/

func Elev_set_motor_direction(dirn int) {
	if dirn == 0 {
		Io_write_analog(MOTOR, 0)
	} else if dirn == 1 {
		Io_clear_bit(MOTORDIR)
		Io_write_analog(MOTOR, MOTOR_SPEED)
	} else if dirn == -1 {
		Io_set_bit(MOTORDIR)
		Io_write_analog(MOTOR, MOTOR_SPEED)
	}
}

/*func Elev_set_motor_direction(dirn config.MotorDirection) {
	if dirn == config.MotorD_stop {
		io_write_analog(MOTOR, 0)
	} else if dirn == config.MotorD_up {
		io_clear_bit(MOTORDIR)
		io_write_analog(MOTOR, MOTOR_SPEED)
	} else if dirn == config.MotorD_down {
		io_set_bit(MOTORDIR)
		io_write_analog(MOTOR, MOTOR_SPEED)
	}
}*/

func Elev_set_button_lamp(button int, floor int, value int) {

	if value == 1 {
		Io_set_bit(lamp_channel_matrix[floor][button])
	} else {
		Io_clear_bit(lamp_channel_matrix[floor][button])
	}
}

func Elev_set_floor_indicator(floor int) {

	// Binary encoding. One light must always be on.
	if floor&0x02 != 0 {
		Io_set_bit(LIGHT_FLOOR_IND1)
	} else {
		Io_clear_bit(LIGHT_FLOOR_IND1)
	}

	if floor&0x01 != 0 {
		Io_set_bit(LIGHT_FLOOR_IND2)
	} else {
		Io_clear_bit(LIGHT_FLOOR_IND2)
	}
}

func Elev_set_door_open_lamp(value int) {
	if value != 0 {
		Io_set_bit(LIGHT_DOOR_OPEN)
	} else {
		Io_clear_bit(LIGHT_DOOR_OPEN)
	}
}

func Elev_set_stop_lamp(value int) {
	if value != 0 {
		Io_set_bit(LIGHT_STOP)
	} else {
		Io_clear_bit(LIGHT_STOP)
	}
}

func Elev_get_button_signal(button config.ElevButtonType, floor int) int {

	return io_read_bit(Button_channel_matrix[floor][button])
}

func Elev_get_floor_sensor_signal() int {
	if io_read_bit(SENSOR_FLOOR1) != 0 {
		return 0
	} else if io_read_bit(SENSOR_FLOOR2) != 0 {
		return 1
	} else if io_read_bit(SENSOR_FLOOR3) != 0 {
		return 2
	} else if io_read_bit(SENSOR_FLOOR4) != 0 {
		return 3
	} else {
		return -1
	}
}

func Elev_get_stop_signal() int {
	return io_read_bit(STOP)
}

/*
func elev_get_obstruction_signal() int {
	return io_read_bit(OBSTRUCTION)
}

type Message stuct{
	Id string
	Msg int
}
*/

/*func Elev_button_pushed(buttonEvents chan<- config.ButtonEvent) {
	for {
		for floor := 0; floor < config.N_FLOORS; floor++ {
			if Elev_get_button_signal(config.Button_Up, floor) == 1 {
				buttonEvents <- config.ButtonEvent{floor, 0}
			} else if Elev_get_button_signal(config.Button_Down, floor) == 1 {
				buttonEvents <- config.ButtonEvent{floor, 1}
			} else if Elev_get_button_signal(config.Button_Command, floor) == 1 {
				buttonEvents <- config.ButtonEvent{floor, 2}
			}
			/*for button := 0; button < config.N_BUTTONS; button++ {
				if Elev_get_button_signal(button, floor) == 1 {
					buttonEvents <- config.ButtonEvent{floor, button}
				}
			}
		}
		time.Sleep(30 * time.Millisecond)
	}

}*/
