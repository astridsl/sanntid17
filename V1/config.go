package V1

const N_FLOORS = 4
const N_BUTTONS = 3

const (
	Button_Up int = iota
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

type ButtonEvent struct {
	Floor  int
	Button int
}

type Status struct {
	LastFloor    int
	Dir          MotorDirection
	CurrentState State
}

type Message struct {
	Source   int
	Category int
	Button   int
	Floor    int
	Status   Status
	Target   int
}

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
