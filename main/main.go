package main

import (
	"../Network"
	//"../Statemachine"
	"../driver"
	//"fmt"
)

func main() {
	//Initializing elevator
	//driver.Elev_init()

	//Initializing statemachine
	//statemachine.initialize()

	//fmt.Println("Hello Github!")

	//Statemachine.initialize()

	Network.Network()

	//driver.Elev_init()

	//driver.Elev_set_motor_direction(-1)

	//Initializing (IKKE INKLUDERT: INIT AV HEIS-HARDWARE-DRIVER OG FSM-GOROUTINE --> MÅ IMPLEMENTERES!!)

	Network.Initialize_network()

	/*
		//Cases for sending
		for {
			select {
			case a := <-buttonPresses:
				if a.Button == B_Cab {
					fsm_tx_newRequests <- Request{a.Floor, a.Button, id, RS_New}
					elev_driver.SetButtonLamp(int(a.Button), a.Floor, 1)
				} else {
					bestId := Delegate(elevators, aliveElevators, a.Floor, a.Button)
					net_tx_requests <- Request{a.Floor, a.Button, bestId, RS_New}
				}

			case a := <-fsm_rx_states:

			case a := <-fsm_rx_errs:

			case a := <-fsm_rx_completedRequests:

			case a := <-net_rx_requests:
				switch a.Status {
				case RS_New:
					if a.Owner == id {
						fsm_tx_newRequests <- a
						net_tx_requests <- Request{a.Floor, a.Button, id, RS_Confirmed}
						elev_driver.SetButtonLamp(int(a.Button), a.Floor, 1)
					}
				case RS_Confirmed:
					elev_driver.SetButtonLamp(int(a.Button), a.Floor, 1)
				case RS_Completed:
					elev_driver.SetButtonLamp(int(a.Button), a.Floor, 0)
				}

			case a := <-net_rx_states:

			case a := <-net_rx_cabRequestRestore:

			case a := <-peer_rx_updates:
	*/
	for {
		select {
		case a := <-Network.Peer_rx_stop_button_pressed:
			driver.Elev_set_stop_lamp(1)

		}

	}

}


test := make(chan )