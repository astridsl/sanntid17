package Network

import (
 	"fmt"
	"log"
	"net"
	"time"
)


func check_for_error(err error){
	if err != nil{
		log.Fatal(err)
	}
}


func UDP1(){
	//set up send-socket
	local_addr, _ := net.ResolveUDPAddr("udp", "")
	remote_addr, _ := net.ResolveUDPAddr("udp", "129.241.187.43:20010")
	socket_send, err := net.DialUDP("udp", local_addr, remote_addr)

	check_for_error(err)

	//set up listen socket
	port, _ := net.ResolveUDPAddr("udp", ":20010")
	socket_listen, err := net.ListenUDP("udp", port)

	check_for_error(err)

	//closing sockets
	defer socket_listen.Close()
	defer socket_send.Close()

	for{
		//sending message
		message := "halloa"
		socket_send.Write([]byte(message))

		//listening to message
		buffer := make([]byte, 1024)
		n, addr, err := socket_listen.ReadFromUDP(buffer[:])
		check_for_error(err)
		//fmt.Println("length : ", length)
		fmt.Println("addr : ", addr)
		fmt.Println(string(buffer[:n]))
		time.Sleep(1*time.Second)

	}



}