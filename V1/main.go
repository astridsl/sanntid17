package V1 

import{
	//Packages and directory 
}

func main() {

	//Make neccesary channels 
	ButtonEvents := make(chan def.ButtonEvent)
	FloorArrivalEvents := make(chan int)
	FloorReached := make(chan int)
	OrderComplete := make(chan int,1) 
	NewOrder := make(chan bool)
	UpdateStatus := make(chan bool, 5) 
	ToNetwork := make(chan def.Message, 100) 
	FromNetwork := make(chan def.Message, 100) 

	fmt.Println("The system is up and running!") //Skal vi ha med denne?? 
	floor, err := driver.Init()  //floor og err er returverdier fra init()

	//Check if error during startUp 
	if err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}

	 am := V1.AssignMaster() //Må se på denne!! 
	 V1.Init(floor, FloorReached, OrderComplete, NewOrder, UpdateStatus) //Må endre ift directory 
	 fmt.Println("Initialised at floor: ", floor) //Print-funksjon... 

	 go elev.ButtonPoller(ButtonEvents) 		//Directory and parameters... 
	 go elev.FloorSensorPoller(FloorArrivalEvents)
	 go network.NetworkManager(ToNetwork, FromNetwork)
	 go safeKill() //Overkill? 
	 queue.Init(NewOrder) //Sjekk denne 

	for {
		select {
		case be := <-ButtonEvents: //Parameter b 
			am.SendNewOrderEvent(be, ToNetwork)

		case fr := <-FloorArrivalEvents: //Parameter f 
			FloorReached <- fr

		case orderCompleteFloor := <-OrderCompleted:
			am.SendOrderCompleteMessage(orderCompleteFloor, ToNetwork)

		case <-UpdateStatus:
			am.SendUpdateStatusMessage(ToNetwork)

		case message := <-FromNetwork:
			switch message.Category {
			case def.ElevAddedToNetwork:
				am.AddNewElev(message, ToNetwork)

			case def.ElevRemovedFromNetwork:
				am.RemoveDeadElev(message.Source, ToNetwork)

			case def.NewOrder:
				am.RegisterNewExtOrder(message, ToNetwork)

			case def.OrderComplete:
				am.FinishedExtOrder(message.Floor)

			case def.UpdateElevStatus:
				am.UpdateElevators(message)
			}
		}
	}																												

}

// safeKill turns the motor off if the program is killed with CTRL+C 

func safeKill() { //Vurdere denne... 
	var c = make(chan os.Signal) //Endre c paramterer... 
	signal.Notify(c, os.Interrupt) //Sjekk notify- funksjon  
	<-c 
	elev.SetMotorDirection(def.MotorD_stop) //directory  
	for floor := 0; floor < def.N_FLOORS; floor++ {
		for button := 0; button < def.N_BUTTONS; button++ {
			elev.SetButtonLight(floor, button, false) //directory 
		}
	}
	elev.SetDoorOpenLight(false) //Sjekk directory på denne 
	fmt.Printf("Program is terminated") //Print- funksjon... 
	os.Exit(1)
}





