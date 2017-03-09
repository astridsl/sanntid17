package queue

//NB: MÅ HUSKE Å ENDRE DE VARIABLENE SOM HØRER TIL config (HAR def. FORANN). KAN VÆRE AT VI MÅ TA BORT def.

import (

)

//Defining variables
var newOrder chan bool

//Local queue --> array med oversikt over alle ordrene som er gitt til denne heisen:
//4 etasjer, 3 type knapper
var localQueue = [def.N_FLOORS][def.N_BUTTONS]int{
	[def.N_BUTTONS]int{0, 0, 0},
	[def.N_BUTTONS]int{0, 0, 0},
	[def.N_BUTTONS]int{0, 0, 0},
	[def.N_BUTTONS]int{0, 0, 0},
}

func Init(tempNewOrder chan bool) {
	newOrder = tempNewOrder //Get new order
	readFromFile() //Get internal backup
	//Set button lamp
	for f := 0; f < def.N_FLOORS; f++ {
		//Check if button is selected
		if localQueue[f][def.Button_Command]==1{
			elev.SetButtonLight(i, def.Button_Command, true) //!!!BRUKT driver.go
			newOrder <- true
		}
		//NY KODE! NB: KAN VÆRE AT DENNE MÅ ENDRES/FJERNES
		if localQueue[f][def.Button_Up]==1{
			elev.SetButtonLight(i, def.Button_Up, true) //!!!BRUKT driver.go
			newOrder <- true
		}
		if localQueue[f][def.Button_Down]==1{
			elev.SetButtonLight(f, def.Button_Down, true) //!!!BRUKT driver.go
			newOrder <- true
		}
		//NY KODE FERDIG!
	}

}

//old: AddToLocalQueue()
func AddOrderToLocalQueue(floor int, button int){
	localQueue[floor][button] = 1
	driver.SetButtonLight(floor,button,true)
	newOrder <- true
	writeToFile() //Update backup
}

//old: RemoveLocalOrderAtFloor()
func DelLocalOrdersAtFloorFromQueue(floor int){
	for b := 0; b < def.N_BUTTONS; b++ {
		localQueue[floor][b]=0; //Delete all local orders at a given floor.
		elev.SetButtonLight(floor,b,false)//Turn off elevator light indecating that there are no orders on this floor.
	}
	writeToFile() //Update backup.

}

//old: ChooseDirection()
func ChooseElevDirection(floor int, direction def.MotorDirection) (def.MotorDirection, def.MotorDirection) {
	switch direction {
	case def.MotorD_down:
		if existsOrdersBelow(floor)==true && floor > 0 {
			return def.MotorD_down, def.MotorD_down
		}
		else if existsOrdersAbove(floor)==true && floor<def.N_FLOORS-1 {
			return def.MotorD_up, def.MotorD_up
		}
		else {
			return def.MotorD_stop, def.MotorD_down
		}

	case def.MotorD_up:
		if existsOrdersAbove(floor)==true && floor<def.N_FLOORS-1 {
			return def.MotorD_up, def.MotorD_up
		}
		else if existsOrdersBelow(floor)==true && floor>0 {
			return def.MotorD_down, def.MotorD_down
		}
		else {
			return def.MotorD_stop, def.MotorD_up
		}

	default:
		return def.MotorD_stop, def.MotorD_up
	}
}

//old ShouldStop()
func ElevShouldStop(floor int, direction def.MotorDirection) bool {
	switch direction {
	case def.MotorD_down:
		if isOrderAtFloor(floor) || floor==0 {
			return true
		}
		else {
			return false
		}

	case def.MotorD_up:
		if isOrderAtFloor(floor) || floor==def.N_FLOORS-1 {
			return true
		}
		else {
			return false
		}

	default:
		return true

	}

}

//old isOrderAt()
func isOrderAtFloor(floor int) bool {
	for b := 0; b < def.N_BUTTONS; b++ {
		if localQueue[floor][b]==1 {
			return true
		}
	}
	return false
}

//old: ordersAboveExist()
func existsOrdersAbove(floor int) bool {
	if floor < def.N_FLOORS-1 {
		for f := floor+1; f < def.N_FLOORS; f++ {
			for b := 0; b < def.N_BUTTONS; b++ {
				if localQueue[f][b]==1 {
					return true
				}
			}
		}
	}
	else {
		return false
	}
	return false
}

//old: ordersBelowExist()
func existsOrdersBelow(floor int) bool {
	if floor > 0 {
		for f := floor-1; f > -1 ; f-- { 
			for b := 0; b < def.N_BUTTONS; b++ {
				if localQueue[f][b]==1 {
					return true
				}
			}
		}
	}
	else {
		return false
	}
	return false
}

//????????????????
//func isOrderAt() {
	
//}





//Read/get information from backup
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


