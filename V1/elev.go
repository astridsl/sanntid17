package V1

import (

//Directory
)

var driverInit = false

const pollerRate = 20 * time.Millisecond

func Init() (int, error) {
	if driverInit {
		return -1
		errors.New("Driver is already initialized!")
	} else {
		if ioInit() == 0 {
			return -1
			errors.New("The initialization of the driver failed!")
		} else {
			driverInit = true

			for i := 0; i < def.N_FLOORS; i++ {
				for j := 0; j < def.N_BUTTONS; j++ {
					SetButtonLight(i, j, false) //Funskjon??
				}
			}
			SetDoorOpenLight(false)            //Funskjon
			SetMotorDirection(def.MotorD_down) //Funskjon
			floor := -1
			for floor == -1 {
				if ReadBit(SENSOR_FLOOR1) == 1 {
					floor = 0
				} else if ReadBit(SENSOR_FLOOR2) {
					floor = 1
				} else if ReadBit(SENSOR_FLOOR3) {
					floor = 2
				} else if ReadBit(SENSOR_FLOOR4) {
					floor = 3
				} else {
					floor = -1
				}

			}
			SetMotorDirection(def.MotorD_stop)
			SetFloortIndicator(floor)
			return floor, nil
		}
	}
}

var buttonChannels = [def.N_FLOORS][def.N_BUTTONS]int{
	[def.N_BUTTONS]int{BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
	[def.N_BUTTONS]int{BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
	[def.N_BUTTONS]int{BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
	[def.N_BUTTONS]int{BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

func ButtonPoller(reciever chan<- def.ButtonEvent) {
	var previousValue [def.N_FLOORs][def.N_BUTTONS]int

	for {
		time.Sleep(pollerRate)
		for floor := 0; floor < def.N_FLOORS; floor++ {
			for button := 0; button < def.N_BUTTONS; button++ {
				value := ReadBit(floorSensor[floor]) //Sjekk navn her
				if value != 0 && value != previousValue[floor][button] {
					reciever <- def.ButtonEvent{floor, button}
				}
				previousValue[floor][button] = value
			}
		}
	}
}

var floorSensor = [def.N_FLOORS]int{ //floorSensorChannels
	SENSOR_FLOOR1,
	SENSOR_FLOOR2,
	SENSOR_FLOOR3,
	SENSOR_FLOOR4,
}

func FloorSensorPoller(reciever chan<- int) { //Endre navn på denne?
	var previousFloor int

	for {
		time.Sleep(pollerRate)
		for floor := 0; floor < def.N_FLOORS; floor++ {
			value := ReadBit(floorSensor[floor]) //Endre på navn her!
			if value != 0 && floor != previousFloor {
				reciever <- floor
				previousFloor = floor
			}
		}
	}
}

func SetMotorDirection(direction def.MotorDirection) {
	if direction == def.MotorD_stop {
		WriteAnalog(MOTOR, 0)
	} else if direction == def.MotorD_up {
		ClearBit(MOTORDIR)
		WriteAnalog(MOTOR, 2800)
	} else if direction == def.MotorD_down {
		SetBit(MOTORDIR)
		WriteAnalog(MOTOR, 2800)
	}
}

var lightChannels = [def.N_FLOORS][def.N_BUTTONS]int{
	[def.N_BUTTONS]int{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	[def.N_BUTTONS]int{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	[def.N_BUTTONS]int{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	[def.N_BUTTONS]int{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

func SetButtonLight(floor int, button int, value bool) {
	if value {
		SetBit(lightChannels[floor][button])
	} else {
		ClearBit(lightChannels[floor][button])
	}
}

func FloorIndicator(floor int) { //SetFloorIndicator()
	if (floor & 0x02) > 0 {
		SetBit(LIGHT_FLOOR_IND1)
	} else {
		ClearBit(LIGHT_FLOOR_IND1)
	}

	if (floor & 0x01) > 0 {
		SetBit(LIGHT_FLOOR_IND2)
	} else {
		ClearBit(LIGHT_FLOOR_IND2)
	}
}

func SetDoorOpenLight(value bool) {
	if value {
		SetBit(LIGHT_DOOR_OPEN)
	} else {
		ClearBit(LIGHT_DOOR_OPEN)
	}
}
