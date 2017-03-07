package Network

import (
	"./network/bcast"
	"./network/localip"
	"./network/peers"
	"flag"
	"fmt"
	"os"
	"time"
)

// We define some custom struct to send over the network.
// Note that all members we want to transmit must be public. Any private members
//  will be received as zero-values.
type HelloMsg struct {
	Message string
	Iter    int
}

func Network() {
	// Our id can be anything. Here we pass it on the command line, using
	//  `go run main.go -id=our_id`
	var id string
	flag.StringVar(&id, "id", "", "id of this peer")
	flag.Parse()

	// ... or alternatively, we can use the local IP address.
	// (But since we can run multiple programs on the same PC, we also append the
	//  process ID)
	if id == "" {
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		id = fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())
	}

	// We make a channel for receiving updates on the id's of the peers that are
	//  alive on the network
	peerUpdateCh := make(chan peers.PeerUpdate)
	// We can disable/enable the transmitter after it has been started.
	// This could be used to signal that we are somehow "unavailable".
	peerTxEnable := make(chan bool)
	go peers.Transmitter(15647, id, peerTxEnable)
	go peers.Receiver(15647, peerUpdateCh)

	// We make channels for sending and receiving our custom data types
	helloTx := make(chan HelloMsg)
	helloRx := make(chan HelloMsg)
	// ... and start the transmitter/receiver pair on some port
	// These functions can take any number of channels! It is also possible to
	//  start multiple transmitters/receivers on the same port.
	go bcast.Transmitter(16569, helloTx)
	go bcast.Receiver(16569, helloRx)

	// The example message. We just send one of these every second.
	go func() {
		helloMsg := HelloMsg{"Hello from " + id, 0}
		for {
			helloMsg.Iter++
			helloTx <- helloMsg
			time.Sleep(1 * time.Second)
		}
	}()

	fmt.Println("Started")
	for {
		select {
		case p := <-peerUpdateCh:
			fmt.Printf("Peer update:\n")
			fmt.Printf("  Peers:    %q\n", p.Peers)
			fmt.Printf("  New:      %q\n", p.New)
			fmt.Printf("  Lost:     %q\n", p.Lost)

		case a := <-helloRx:
			fmt.Printf("Received: %#v\n", a)
		}
	}
}

func Initialize_network() {
	/*
		//Initializing network
		peer_rx_updates := make(chan peers.PeerUpdate)
		peer_tx_enable := make(chan bool)
		go peers.Transmitter(28455, id, peer_tx_enable)
		go peers.Receiver(28455, peer_rx_updates)

		net_tx_requests := make(chan Request)
		net_rx_requests := make(chan Request)
		net_tx_states := make(chan ElevatorUpdate)
		net_rx_states := make(chan ElevatorUpdate)
		net_tx_cabRequestRestore := make(chan CabRequestRestore)
		net_rx_cabRequestRestore := make(chan CabRequestRestore)
		go bcast.Transmitter(28456, net_tx_requests, net_tx_states, net_tx_cabRequestRestore)
		go bcast.Receiver(28456, net_rx_requests, net_rx_states, net_rx_cabRequestRestore)
	*/
	StopButtonPressed := make(chan int)
	Peer_rx_stop_button_pressed := make(chan StopButtonPressed)
	go bcast.Transmitter(28456, Peer_rx_stop_button_pressed)
	go bcast.Receiver(28456, Peer_rx_stop_button_pressed)
}
