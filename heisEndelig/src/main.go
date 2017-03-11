package src

import {
	"./Network"
	"./config"
}

func main() {
	
	elev.Elev_init()

	if Elev_init()==nil{
		fmt.println("Initialization failed")
	}

	floor = elev_get_floor_sensor_signal()//Get the floor elevator is at after initializing

	go Network.main()//Setter i gang network og sjekker etter og bestemmer hele tiden  master

	ButtonEvents := make(chan config.ButtonEvent) //Channel som sender external order til nettverk 
	OrderCompleted := make(chan bool)
	ArrivedAtFloor := make(chan int)
	NewOrder := make(chan bool) 


	statemachine.initialize(NewOrder, ArrivedAtFloor, OrderCompleted, floor) 

	go elev.Elev_button_pushed(ButtonEvents)

	for {
		select {
		case <- ButtonEvents:
			queue.AddOrderToLocalQueue(ButtonEvents.floor, ButtonEvents.button)
		
		//case <- OrderCompleted: 
			

		case <- ArrivedAtFloor: 



		case <- NewOrder: 

		}
	}

	//Cost




	/*for {
		if Network.isMaster {
			//Heis er master --> deleger ordre

		} else {
			//Heis er slave --> do something

		}
	}*/

	//HUSK Å INITIALISERE driver MED elev.init --> LAG ERROR-FUNKSJON HVIS DET SKJER NOE FEIL

}



