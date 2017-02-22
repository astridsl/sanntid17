#ifndef STATEMACHINE_H_
#define STATEMACHINE_H_


extern int motorDirection;
extern int nextFloor;
extern int lastFloor;
extern int lastFloorStoped;


//Initialiserer heisen. Om heisen ikke står i en etasje kjøres den ned til nærmeste etasje.
void initialize(void);

//Setter hvilken tilstand heisen er. Utfører valgte spesifikasjoner for valgt tilstand og trigger.
void driveElevator(void);

//Returnerer 1 hvis en knapp er trykket inn, og 0 ellers.
int evButtonSignalPressed(void);

//Returnerer 1 hvis heisen står i nextFloor og etasjen er bestilt, returnerer -1 hvis heisen står i nextFloor og etasjen ikke er bestilt. Returnerer 0 ellers.
int evAtSelectedFloor(void);

//Returnerer forige kjøreretning heisen hadde eller neste dersom heisen er i 1./4. etasje, 1 = opp og -1 = ned. Samt setter mottoren til opp, ned eller stopp.
int motor_direction(void);

//Setter nødvendige variabler.
void eventManager(void);

//Heisens tilstander. 
typedef enum state { 
    	state_idle,
	state_moving,
	state_door_open,
	state_emergency_stop,
	state_undefined

}state_t;

state_t current_state;


#endif /*STATEMACHINE_H_*/
