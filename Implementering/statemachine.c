#include "statemachine.h"
#include "elev.h"
#include "queue.h"
#include "timer.h"
#include <stdio.h>


#include <assert.h>
#include <stdlib.h>


int motorDirection;
int nextFloor;
int lastFloor;
int lastFloorStoped;
static int currentFloor;
static int buttonPressed;
static int atSelectedFloor;
static int timeout;
static int stopSignalPressed;


void  initialize(void){
	queue_del_queue();
	elev_set_floor_indicator(0);
	currentFloor = elev_get_floor_sensor_signal();
	if (currentFloor == -1){
		elev_set_motor_direction(DIRN_DOWN);
		while(currentFloor == -1){
			currentFloor = elev_get_floor_sensor_signal();
		}
		elev_set_motor_direction(DIRN_STOP);
	}
	nextFloor = currentFloor;
	lastFloor = currentFloor;
	lastFloorStoped = currentFloor;
	elev_set_floor_indicator(lastFloor);
	motorDirection = -1;
	current_state = state_idle;
}


int evButtonSignalPressed(void){
	int floor;
	for(floor=0; floor<N_FLOORS; floor++){
		if((floor!=3) && (elev_get_button_signal(BUTTON_CALL_UP, floor) == 1)){
			queue_add_to_queue(floor, BUTTON_CALL_UP);
			return 1;
		}else if((floor!=0) && (elev_get_button_signal(BUTTON_CALL_DOWN,floor)==1)){
			queue_add_to_queue(floor, BUTTON_CALL_DOWN);
			return 1;
		}else if(elev_get_button_signal(BUTTON_COMMAND,floor)==1){
			queue_add_to_queue(floor, BUTTON_COMMAND);
			return 1;
		}
	}
	return 0;
}


int evAtSelectedFloor(void){
	int button;
	for (button = 0; button < N_BUTTONS; button++){
		if((elev_get_floor_sensor_signal() == nextFloor) && (get_button_pressed(nextFloor,button)==1)){
			return 1;
		}
	}
	if (elev_get_floor_sensor_signal() == nextFloor){
		return -1;
	}
	return 0;
}


int motor_direction(void){
	if (lastFloor == nextFloor || nextFloor == -1){ 
		elev_set_motor_direction(DIRN_STOP);
		if (lastFloor == N_FLOORS-1){
			return -1;
		}else if (lastFloor == 0){
			return 1;
		}
		return motorDirection;
	}else if (lastFloor < nextFloor && current_state != state_door_open ){
		elev_set_motor_direction(DIRN_UP);
		return 1;
	}else if (lastFloor > nextFloor && current_state != state_door_open){
		elev_set_motor_direction(DIRN_DOWN);
		return -1;
	}
	return motorDirection;
}


void eventManager(void){
	stopSignalPressed = elev_get_stop_signal();
	if (stopSignalPressed == 0){
		buttonPressed = evButtonSignalPressed();
		atSelectedFloor = evAtSelectedFloor();
		currentFloor = elev_get_floor_sensor_signal();
		motorDirection = motor_direction();
		nextFloor = queue_get_next_floor();
		timeout = evIsTimeout();
		if (currentFloor != -1){
			lastFloor = currentFloor;
			elev_set_floor_indicator(lastFloor);
		}
		//Følgende kan bare gjelde dersom heisen har vært stoppet. Sjekker om neste etasje er den forige etasjen heisen kjørte forbi slik at man skal kunne kjøre tilbake til denne etasjen dersom den bestilles. 
		if (nextFloor == lastFloor && currentFloor == -1 ) {
			if (motorDirection == 1){
				lastFloor++;
			}else{
				lastFloor--;
			}
		}
	}else {
		elev_set_stop_lamp(1);
		elev_set_motor_direction(0);
		queue_del_queue();
		current_state = state_emergency_stop;
		nextFloor = -1; //Indikerer at heisen skal stå stille helt til den får en ny bestilling.
	}
}

void driveElevator(void){
	switch(current_state){
		case state_idle:
			if(atSelectedFloor == 1){
				timer_start();
				elev_set_door_open_lamp(1);
				queue_del_one_floor();
				lastFloorStoped = lastFloor;
				current_state = state_door_open;
			}else if (atSelectedFloor == 0){
				current_state = state_moving;
			}
			break;
		case state_moving:
			if (atSelectedFloor == 1){
				timer_start();
				elev_set_door_open_lamp(1);
				queue_del_one_floor();
				lastFloorStoped = lastFloor;
				current_state = state_door_open;
			}
			break;
		case state_door_open:
			if (timeout == 1){
				elev_set_door_open_lamp(0);
				current_state = state_idle;
			}else if (buttonPressed == 1 && nextFloor == lastFloorStoped){
				timer_start();
				queue_del_one_floor();
			}
			break;
		case state_emergency_stop:
			if (stopSignalPressed == 0 && currentFloor != -1){
				elev_set_stop_lamp(0);
				elev_set_door_open_lamp(0);
				current_state = state_idle;
			}else if (currentFloor != -1){
				elev_set_door_open_lamp(1);
			}else if (stopSignalPressed == 0){
				elev_set_stop_lamp(0);
				current_state = state_undefined;
			}
			break;
		case state_undefined:
			if (buttonPressed == 1){
				current_state = state_moving;
			}
			break;
	}
}
			



