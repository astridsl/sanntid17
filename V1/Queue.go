package queue

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
	for i := 0; i < def.N_FLOORS; i++ {
		//Check if button is selected
		if localQueue[i][def.Button_Command]==1{
			driver.SetButtonLight(i, def.Button_Command, true) //!!!BRUKT driver.go
			newOrder <- true
		}
		//NY KODE! NB: KAN VÆRE AT DENNE MÅ ENDRES/FJERNES
		if localQueue[i][def.Button_Up]==1{
			driver.SetButtonLight(i, def.Button_Up, true) //!!!BRUKT driver.go
			newOrder <- true
		}
		if localQueue[i][def.Button_Down]==1{
			driver.SetButtonLight(i, def.Button_Down, true) //!!!BRUKT driver.go
			newOrder <- true
		}
		//NY KODE FERDIG!
	}

}

func AddOrderToLocalQueue(floor int, button int){
	localQueue[floor][button] = 1
	driver.SetButtonLight(floor,button,true)
	newOrder <- true
	writeToFile() //Update backup
}

func DelLocalOrderFromQueueAtFloor(floor int){
	

}



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

