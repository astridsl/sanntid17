package Order

//Matrise for å håndtere køen.
var buttonMatrix[N_FLOORS][N_BUTTONS] int = {{0}}
 


func get_button_pressed(int floor, int button)int{
	return buttonMatrix[floor][button]
}


func queue_del_queue(){
	floor int
	button int 
	for floor=0; floor<N_FLOORS; floor++{
		for button=0; button<N_BUTTONS; button++{
			buttonMatrix[floor][button]=0
			if floor!=0 && button == 1{
				elev_set_button_lamp(BUTTON_CALL_DOWN, floor, 0)
			}else if floor!=N_FLOORS-1 && button == 0{
				elev_set_button_lamp(BUTTON_CALL_UP, floor, 0);
			}else if button == 2{
				elev_set_button_lamp(BUTTON_COMMAND, floor, 0);
			}
		}
	}
}


func queue_del_one_floor(){
	button int
	for button=0; button<N_BUTTONS; button++{
		buttonMatrix[lastFloor][button]=0
		if lastFloor!=0 && button == 1{
			elev_set_button_lamp(BUTTON_CALL_DOWN, lastFloor, 0)
		}else if lastFloor!=N_FLOORS-1 && button == 0{
			elev_set_button_lamp(BUTTON_CALL_UP, lastFloor, 0)
		}else if button == 2{
			elev_set_button_lamp(BUTTON_COMMAND, lastFloor, 0)
		}
	}
}


func queue_add_to_queue(int floor, int button){
	buttonMatrix[floor][button]=1
	elev_set_button_lamp(button, floor, 1)
}


func queue_get_next_floor()int{
	int floor
	//Sjekker om det er en bestilling for den etasjen heisen står stille i.
	if elev_get_floor_sensor_signal() == lastFloorStoped{
		if (buttonMatrix[lastFloorStoped][BUTTON_COMMAND] == 1 || buttonMatrix[lastFloorStoped][BUTTON_CALL_DOWN] == 1 || buttonMatrix[lastFloorStoped][BUTTON_CALL_UP] == 1){
			return lastFloorStoped
		}
	}
	//Sjekker først om det er en bestilling under heisen før det sjekkes for etasjene over.
	if motorDirection == -1{
		for floor = lastFloor-1; floor > 0; floor--{
			if (buttonMatrix[floor][BUTTON_COMMAND] == 1 || buttonMatrix[floor][BUTTON_CALL_DOWN] == 1){
				return floor
			}	
		}
		for floor = 0; floor < N_FLOORS-1; floor++{
			if buttonMatrix[floor][BUTTON_COMMAND] == 1 || buttonMatrix[floor][BUTTON_CALL_UP] == 1{
				return floor
			}	
		}
		for floor = lastFloor-1; floor < N_FLOORS; floor++{
			if (buttonMatrix[floor][BUTTON_COMMAND] == 1 || buttonMatrix[floor][BUTTON_CALL_DOWN] == 1){
				return floor
			}	
		}
	}
	//Sjekker først om det er en bestilling over heisen før det sjekkes for etasjene under.
	else if motorDirection == 1{
		for floor = lastFloor+1; floor < N_FLOORS-1; floor++{
			if buttonMatrix[floor][BUTTON_COMMAND] == 1 || buttonMatrix[floor][BUTTON_CALL_UP] == 1{
				return floor
			}	
		}
		for floor = N_FLOORS-1; floor > 0; floor--{
			if buttonMatrix[floor][BUTTON_COMMAND] == 1 || buttonMatrix[floor][BUTTON_CALL_DOWN] == 1{
				return floor
			}	
		}
		for floor = lastFloor+1; floor >= 0; floor--{
			if buttonMatrix[floor][BUTTON_COMMAND] == 1 || buttonMatrix[floor][BUTTON_CALL_UP] == 1{
				return floor
			}	
		}
		
	}
	return nextFloor
}
