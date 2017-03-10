package network

import (
	def "config"
	"net"
	"time"
	"udp"
)

var elev_timers map[int]*time.Timer

func NetworkManager(elevToNetwork chan def.Message, networkToElev chan def.Message) {
	addr, _ := net.InterfaceAddrs()
	myID := int(addr[1].String()[12]-'0')*100 + int(addr[1].String()[13]-'0')*10 + int(addr[1].String()[14]-'0')

	UDPsend := make(chan def.Message, 100)
	UDPreceive := make(chan def.Message, 100)

	go broadcastIP(myID, UDPsend)
	go udp.UDPListener(UDPreceive)
	go udp.UDPSender(UDPsend)

	elev_timers = make(map[int]*time.Timer)

	for {
		select {
		case message := <-UDPreceive:
			_, present := elev_timers[message.Source]

			if message.Category == def.IP {
				if message.Source != myID {
					if present {
						elev_timers[message.Source].Reset(2 * time.Second)
					} else {
						elev_timers[message.Source] = time.AfterFunc(2*time.Second, func() { remove_elev(message.Source, elevToNetwork) })
						networkToElev <- def.Message{Source: message.Source, Category: def.ElevAddedToNetwork}
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

func broadcastIP(ID int, UDPsend chan def.Message) {
	for {
		UDPsend <- def.Message{Source: ID, Category: def.IP}
		time.Sleep(100 * time.Millisecond)
	}
}

func remove_elev(ID int, networkToElev chan def.Message) {
	delete(elev_timers, ID)
	networkToElev <- def.Message{Source: ID, Category: def.ElevRemovedFromNetwork}
}
