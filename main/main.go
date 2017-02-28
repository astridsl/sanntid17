package main

import (
	//"./Interface/driver"
	"../Network"
	//"../Statemachine"
	"fmt"
)

func main() {
	//fmt.Println("Hello Github!")

	//Statemachine.initialize()
	
	//Network.Network()
	
	//driver.Elev_init()

	//driver.Elev_set_motor_direction(-1)
	//Network.NetworkManager("Hello", "HelloBack")
	
	//Initializing (IKKE INKLUDERT: INIT AV HEIS-HARDWARE-DRIVER OG FSM-GOROUTINE --> MÅ IMPLEMENTERES!!)
	init()
	
	
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

        }
    }
	


	
	
}



//Initializing (IKKE INKLUDERT: INIT AV HEIS-HARDWARE-DRIVER OG FSM-GOROUTINE --> MÅ IMPLEMENTERES!!)
func init() {
	peer_rx_updates             := make(chan peers.PeerUpdate)
    peer_tx_enable              := make(chan bool)
    go peers.Transmitter(28455, id, peer_tx_enable)
    go peers.Receiver(   28455,     peer_rx_updates)

    net_tx_requests             := make(chan Request)
    net_rx_requests             := make(chan Request)
    net_tx_states               := make(chan ElevatorUpdate)
    net_rx_states               := make(chan ElevatorUpdate)
    net_tx_cabRequestRestore    := make(chan CabRequestRestore)
    net_rx_cabRequestRestore    := make(chan CabRequestRestore)
    go bcast.Transmitter(28456, net_tx_requests, net_tx_states, net_tx_cabRequestRestore)
    go bcast.Receiver(   28456, net_rx_requests, net_rx_states, net_rx_cabRequestRestore)


}


