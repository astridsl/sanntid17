package main

import (
	//"./Network"
	//"../config"
	//"../driver"
	//"../queue"
	"../statemachine"
	//"fmt"
	"time"
)

func main() {
	//driver.Elev_init()

	statemachine.Initialize()

	for {
		statemachine.EventManager()
		statemachine.DriveElevator()
		time.Sleep(15 * time.Millisecond)
	}

	//TO DO:
	//LAGE ET NETTVERK, BRUKE TO CHANNELS FOR SENDING OG RESCIVEING AV MELDINGER OG KJØRE NETTVERKET SOM EN GO ROUTINE I main.

	/*driver.Elev_init()

	if driver.Elev_init() == 0 {
		fmt.Println("Initialization failed")
	}

	//driver.Elev_set_button_lamp(2, 1, 1)

	fmt.Println("hei")
	/*
		driver.Elev_set_motor_direction(config.MotorD_up)
		//floor := driver.Elev_get_floor_sensor_signal() //Get the floor elevator is at after initializing

		//go Network.main()//Setter i gang network og sjekker etter og bestemmer hele tiden  master

		/*ButtonEvents := make(chan config.ButtonEvent) //Channel som sender external order til nettverk
		//OrderCompleted := make(chan bool)
		ArrivedAtFloor := make(chan int)
		NewOrder := make(chan bool)
	*/

	/*ArrivedAtFloor := make(chan bool)
	NewOrder := make(chan bool)
	floor := driver.Elev_get_floor_sensor_signal()
	//queue.Init()
	statemachine.Initialize(NewOrder, ArrivedAtFloor, floor)

	ButtonEvents := make(chan config.ButtonEvent)
	//NewOrder := make(chan bool)
	//config.ButtonEvent.Floor = 1
	//config.ButtonEvent.Button = 1

	//driver.Elev_button_pushed(ButtonEvents)

	go driver.Elev_button_pushed(ButtonEvents)

	for {
		select {
		case BF := <-ButtonEvents:
			floor := BF.Floor
			button := BF.Button
			driver.Elev_set_button_lamp(button, floor, 1)
			queue.AddOrderToLocalQueue(floor, button)
			NewOrder <- true
			//case <- OrderCompleted:

			//case <-ArrivedAtFloor:

			//case <-NewOrder:
			/*default:
				fmt.Println("do nothing")
			}*/
	//}
	/*
		//Cost

		for {
			if Network.isMaster {
				//Heis er master --> deleger ordre

			} else {
				//Heis er slave --> do something

			}
		}*/

	//}

}
