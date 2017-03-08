package V1

import{
	
	//Directory 
}

func (nm *networkManager) mostSuitable(orderButton int, orderFloor int) int{   //cost- module 
	var cost int 
	minimalCost := 5 //Sjekk ut denne her, får feil her ift til decla
	mostSuitableElevator := nm.myID
	for key := range nm.elevators {
		cost = nm.costFunction(nm.elevators[key], orderButton, orderFloor)
		if cost<minimalCost {
			minimalCost = cost
			mostSuitableElevator = nm.elevators[key]
		} else if cost == minimalCost{
			if nm.elevators[key] > mostSuitableElevator{
				mostSuitableElevator = nm.elevators[key]
			}	
		}
	}
   return mostSuitableElevator 
}


func (nm *networkManager) costFunction(ID int, orderButton int, orderFloor int) int { //cost-module 
	lastFloor := nm.statuses[ID].LastFloor 
	dir := nm.statuses[ID].Dir 
	state := nm.statuses[ID].CurrentState 
	cost := 0 

	if lastFloor == orderFloor && dir == def.MotorD_stop{
		return cost 
	}
	if nm.extOrders[orderFloor][def.Button_Up] ==ID || nm.extOrders[orderFloor][def.Button_Down] == ID {
		return cost 
	}

	if state == 0{
		cost = int(math.Abs(float64(lastFloor - orderFloor)))
	}
	else if lastFloor < orderFloor && dir == def.MotorD_up && orderButton == def.Button_Up || orderFloor < lastFloor && orderButton == def.Button_Down{
		cost = int(math.Abs(float64(orderFloor - lastFloor)))
	} else {
		cost = 5 
	}
	return cost 
}








