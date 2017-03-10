package V1

import{
	//Directory 
}

type networkManager struct{
	myID		int 
	elevators	[]int 
	extOrders	[def.N_FLOORS][2]int 
	masterID	int 
	statuses 	map[int]def.Status   
}


func AssignMaster() networkManager{ 

	var nm networkManager 
	nm.elevators = make([]int, 1) 
	nm.statuses = make(map[int]def.Status) 

	addr, _ := net.InterfaceAddrs()
	nm.myID = int(addr[1].String()[12]-'0')*100 + int(addr[1].String()[13]-'0')*10 + int(addr[1].String()[14]-'0')
	nm.elevators[0] = nm.myID

	myStatus := def.Status{LastFloor: lastFloor, Dir: dir, CurrentState: state}
	nm.statuses[nm.myID] = myStatus 

	nm.masterSelect()
	return nm 
}  


func (nm *networkManager) RemoveElevator(remElevID int, toNetwork chan def.Message){
	for i := range nm.elevators{
		if nm.elevators[i] == remElevID {
			nm.elevators[i] = nm.elevators[len(nm.elevators)-1]
			nm.elevators = nm.elevators[:len(nm.elevators)-1]
			break
		}
	}
			nm.masterSelect()

			for floor := 0; floor < def.N_FLOORS; floor++{
				for button := def.Button_Up; button <= def.Button_Down; button++{
					if nm.extOrders[floor][button] == remElevID{
						nm.extOrders[floor][button] = nm.masterID
						if nm.myID == nm.masterID{
							queue.AddOrderToLocalQueue(floor, button) 
						}
					}
				}
			}
}
		

func (nm *networkManager) AddElevator(message def.Message, toNetwork chan def.Message){
	nm.elevators = append(nm.elevators, message.Source)

	if nm.masterID == nm.myID{
		for floor := 0; floor < def.N_FLOORS; floor++{
			for button :=def.Button_Up; button <= def.Button_Down; button++{
				if nm.extOrders[floor][button] != 0{
					toNetwork <- def.Message{Source: nm.extOrders[floor][button], Category: def.NewOrder, Button: button, Floor: floor, Target: nm.extOrders[floor][button]}

				}
			}
		}
	}
	nm.masterSelect()
}


func (nm *networkManager) newExternalOrder(message def.Message, toNetwork chan def.Message){
	driver.SetButtonLight(message.Floor, message.Button, true)  
	if nm.masterID == nm.myID && message.Target == def.MASTER {
		target := nm.mostSuitable(message.Button, message.Floor)
		toNetwork <- def.Message{Source: nm.myID, Category: def.NewOrder, Button: message.Button, Floor: message.Floor, Target: target}
	} else if message.Source == nm.masterID && message.Target != def.Master {
		nm.extOrders[message.Floor][message.Button] = message.Target 
		if nm.myID == message.Target {
			queue.AddOrderToLocalQueue(message.Floor, message.Button) 
		}
	}
}

func (nm *networkManager) ExternalOrderFinito(floor int){
	for button := def.Button_Up; button <= def.Button_Down; button++{
		nm.extOrders[floor][button] = 0
		driver.SetButtonLight(floor, button, false)
	}
}

func (nm *networkManager) UpdateElevators(message def.Message){
	nm.statuses[message.Source] = message.Status
}

func (nm *networkManager) SendNewOrder(buttonEvent def.ButtonEvent, toNetwork chan def.Message){
	if buttonEvent.Button == def.Button_Command{
		queue.AddOrderToLocalQueue(buttonEvent.Floor, buttonEvent.Button) 
	} else {
		nm.sendOrderMessage(buttonEvent, toNetwork)
	}
}

func (nm *networkManager) sendOrderComplete(floor int, toNetwork chan def.Message){
	nm.ExternalOrderFinito(floor)
	toNetwork <- def.Message{Source: nm.myID, Category: def.OrderComplete, Floor: floor}
}

func (nm *networkManager) sendUpdateStatus(toNetwork chan def.Message){
	newStatus := def.Status{LastFloor: lastFloor, Dir: dir, CurrentState: state}
	toNetwork <- def.Message{Source: nm.myID, Category: def.UpdateElevStatus, Status: newStatus} 
}

func sendOrder(){
	toNetwork <- def.Message{Source: nm.myID, Category: def.NewOrder, Button: b.Button, Floor: b.Floor, Target: def.MASTER} 
}

func (nm *networkManager) masterSelect(){ 
	factor := 256 //Se på denne... 
	for i := range nm.elevators{
		if factor>nm.elevators[i]{
			factor = nm.elevators[i]	
		}
	}
	nm.masterID = factor
	fmt.Println("The new master is ", nm.masterID) //Printfunksjon...

}

