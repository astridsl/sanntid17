package config

const N_FLOORS = 4
const N_BUTTONS = 3

type ElevButtonType int

const (
	Button_Up ElevButtonType = iota
	Button_Down
	Button_Command
)

type MotorDirection int

const (
	MotorD_down MotorDirection = iota - 1
	MotorD_stop
	MotorD_up
)

type State int

const (
	Idle State = iota
	Moving
	DoorOpen
)

//IKKE BRUKT ENDA
//MÅ BRUKE!!!!!!!!! NÅ ER DEN ENDELIG BRUKT OMG OMG!!!!!!!
type ButtonEvent struct {
	Floor  int
	Button int
}

//IKKE BRUKT ENDA
type Status struct {
	LastFloor    int
	Dir          MotorDirection
	CurrentState State
}

//Ikke brukt enda
type Message struct {
	Source   int
	Category int
	Button   int
	Floor    int
	Status   Status
	Target   int
}

//sendMsgCh := make(chan config.Message)
//recMsgCh : = make(chan config.Message)

//sendMsgCh <- Message(-1,)

//IKKE BRUKT ENDA
const (
	// Target constant:
	MASTER int = -1
	// Message category constants:
	ElevAddedToNetwork     = 100
	ElevRemovedFromNetwork = 101
	OrderComplete          = 102
	UpdateElevStatus       = 103
	IP                     = 104
	NewOrder               = 105
)
