
package Network

import (
	"fmt"
	"time"
	"sort"
)


func main(){
	ip := GetLocalIP()
	fmt.Println("Local ip: ", ip)

	go UdpSendAlive()
	peerListCh := make(chan []string)//Liste med alle heisene (ip-adresse)
	go UdpRecvAlive(peerListCh)
	sendMsgCh := make(chan Msg)
	recvMsgCh := make(chan Msg)
	go UdpSendMsg(sendMsgCh)
	go UdpRecvMsg(recvMsgCh)

	tick := time.Tick(1*time.Second)

	isMaster := false
	peers := []string{ip} //Liste med alle ip-adressene til heisene

	for {
		select {
		case <-tick:
			sendMsgCh <- Msg{5, []bool{true, false}, "129.241.187.150"}

		//Velger master
		case peers = <-peerListCh:
			fmt.Println("New peer list: ", peers)
			sort.Strings(peers)
			if len(peers) == 0 {
				fmt.Println("Disconnected, defaulting to master")
				isMaster = true
			} else {
				if peers[0] == ip {
					fmt.Println("We have highest ip, we are master")
					isMaster = true
				} else {
					fmt.Println("We do not have highest ip, we are slave")
					isMaster = false
				}
			}
			fmt.Println("the master is: ", isMaster)
			
		case r := <-recvMsgCh:
			fmt.Println("New msg: ", r)
		}
	}


}

//Ting til meg selv:
//if len(network.peers)==1 --> kun en heis i nettwerket --> trenger ikke å gjøre like mange ting

//if mer enn en heis i nettwerk og master: Håndter ordre (costfunksjon)

//if mer enn en heis i nettverk og slave: gjør noe annet