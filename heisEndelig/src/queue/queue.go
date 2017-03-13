package queue

//NB: MÅ HUSKE Å ENDRE DE VARIABLENE SOM HØRER TIL config (HAR def. FORANN). KAN VÆRE AT VI MÅ TA BORT def.

import (
	"../config"
	"../driver"
)

//Defining variables
//var newOrder chan bool

//Local queue --> array med oversikt over alle ordrene som er gitt til denne heisen:
//4 etasjer, 3 type knapper
var LocalQueue = [config.N_FLOORS][config.N_BUTTONS]int{
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 0},
}

//TA BORT CHANNEL SOM ARGUMENT HER??
func Init() {
	//KODE UTEN BACKUP
	//Sett alle button lys og hele den lokale køen til null
	for floor := 0; floor < config.N_FLOORS; floor++ {
		for button := 0; button < config.N_BUTTONS; button++ {
			driver.Elev_set_button_lamp(button, floor, 1)
			//localQueue[floor][button] == 0
		}
	}

	//KODE NÅR VI HAR BACKUP
	//newOrder = tempNewOrder //Get new order
	//readFromFile()          //Get internal backup
	//Set button lamp
	/*for floor := 0; floor < config.N_FLOORS; floor++ {
		//Check if button is selected
		if localQueue[floor][config.Button_Command] == 1 {
			driver.Elev_set_button_lamp(config.Button_Command, floor, true) //!!!BRUKT driver.go
			//newOrder <- true
		}
		//NY KODE! NB: KAN VÆRE AT DENNE MÅ ENDRES/FJERNES
		if localQueue[f][def.Button_Up]==1{
			driver.SetButtonLight(i, def.Button_Up, true) //!!!BRUKT driver.go
			newOrder <- true
		}
		if localQueue[f][def.Button_Down]==1{
			driver.SetButtonLight(f, def.Button_Down, true) //!!!BRUKT driver.go
			newOrder <- true
		}
		//NY KODE FERDIG!
	}*/

}

//old: AddToLocalQueue()
func AddOrderToLocalQueue(floor int, button int) {
	LocalQueue[floor][button] = 1
	driver.Elev_set_button_lamp(button, floor, 1)
	//newOrder <- true
	//writeToFile() //Update backup
}

//old: RemoveLocalOrderAtFloor()
func DelLocalOrderAtFloor(floor int) { //Funksjonsnavn på endres!
	for b := 0; b < config.N_BUTTONS; b++ {
		LocalQueue[floor][b] = 0                 //Delete all local orders at a given floor.
		driver.Elev_set_button_lamp(b, floor, 0) //Turn off elevator light indecating that there are no orders on this floor.
	}
	//writeToFile() //Update backup.
}

//old: ChooseDirection()
func ChooseElevDirection(floor int, direction config.MotorDirection) (config.MotorDirection, config.MotorDirection) {
	switch direction {
	case config.MotorD_down:
		if existsOrdersBelow(floor) == true && floor > 0 {
			return config.MotorD_down, config.MotorD_down
		} else if existsOrdersAbove(floor) == true && floor < config.N_FLOORS-1 {
			return config.MotorD_up, config.MotorD_down
		} else {
			return config.MotorD_stop, config.MotorD_down
		}

	case config.MotorD_up:
		if existsOrdersAbove(floor) == true && floor < config.N_FLOORS-1 {
			return config.MotorD_up, config.MotorD_up
		} else if existsOrdersBelow(floor) == true && floor > 0 {
			return config.MotorD_down, config.MotorD_up
		} else {
			return config.MotorD_stop, config.MotorD_up
		}

	default:
		return config.MotorD_stop, config.MotorD_up
	}
}

//old ShouldStop()
func ElevShouldStop(floor int, direction config.MotorDirection) bool {
	switch direction {
	case config.MotorD_down:
		if IsOrderAtFloor(floor) || floor == 0 {
			return true
		} else {
			return false
		}

	case config.MotorD_up:
		if IsOrderAtFloor(floor) || floor == config.N_FLOORS-1 {
			return true
		} else {
			return false
		}

	default:
		return true

	}

}

//old isOrderAt()
func IsOrderAtFloor(floor int) bool {
	for b := 0; b < config.N_BUTTONS; b++ {
		if LocalQueue[floor][b] == 1 {
			return true
		}
	}
	return false
}

//old: ordersAboveExist()
func existsOrdersAbove(floor int) bool {
	if floor < config.N_FLOORS-1 {
		for f := floor + 1; f < config.N_FLOORS; f++ {
			for b := 0; b < config.N_BUTTONS; b++ {
				if LocalQueue[f][b] == 1 {
					return true
				}
			}
		}
	} else {
		return false
	}
	return false
}

//old: ordersBelowExist()
func existsOrdersBelow(floor int) bool {
	if floor > 0 {
		for f := floor - 1; f > -1; f-- {
			for b := 0; b < config.N_BUTTONS; b++ {
				if LocalQueue[f][b] == 1 {
					return true
				}
			}
		}
	} else {
		return false
	}
	return false
}

//????????????????
//func isOrderAt() {

//}

//Read/get information from backup
/*
func readFromFile() {
	b, err := ioutil.ReadFile("internalBackup.txt")
	if err != nil {
		writeToFile()
	}

	for i := 0; i < def.N_FLOORS; i++ {
		localQueue[i][def.Button_Command] = int(b[i])
	}
}

//Write to backup
func writeToFile() {
	b := make([]byte, def.N_FLOORS)
	for i := 0; i < def.N_FLOORS; i++ {
		b[i] = byte(localQueue[i][def.Button_Command])
	}

	err := ioutil.WriteFile("internalBackup.txt", b, 0644)
	if err != nil {
		panic(err)
	}
}
*/
