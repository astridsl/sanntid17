package Network

import (
	//def "config"
	"net"
	"time"

	//"udp"
)

var elev_timers map[int]*time.Timer

func NetworkManager(elevToNetwork chan Message, networkToElev chan Message) {
	addr, _ := net.InterfaceAddrs()
	myID := int(addr[1].String()[12]-'0')*100 + int(addr[1].String()[13]-'0')*10 + int(addr[1].String()[14]-'0')

	UDPsend := make(chan Message, 100)
	UDPreceive := make(chan Message, 100)

	go broadcastIP(myID, UDPsend)
	go UDPListener(UDPreceive) //Blir feil her
	go UDPSender(UDPsend)      //Blir feil her

	elev_timers = make(map[int]*time.Timer)

	for {
		select {
		case message := <-UDPreceive:
			_, present := elev_timers[message.Source]

			if message.Category == IP {
				if message.Source != myID {
					if present {
						elev_timers[message.Source].Reset(2 * time.Second)
					} else {
						elev_timers[message.Source] = time.AfterFunc(2*time.Second, func() { remove_elev(message.Source, elevToNetwork) })
						networkToElev <- Message{Source: message.Source, Category: ElevAddedToNetwork}
					}
				}
				break
			}
			networkToElev <- message

		case message := <-elevToNetwork:
			UDPsend <- message
		}
	}
}

func broadcastIP(ID int, UDPsend chan Message) {
	for {
		UDPsend <- Message{Source: ID, Category: IP}
		time.Sleep(100 * time.Millisecond)
	}
}

func remove_elev(ID int, networkToElev chan Message) {
	delete(elev_timers, ID)
	networkToElev <- Message{Source: ID, Category: ElevRemovedFromNetwork}
}
