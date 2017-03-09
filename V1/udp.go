package udp

import (
	def "config"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

const (
	PORT = ":20909"
)

func UDPSender(channel chan def.Message) {
	broadcastAddr := []string{"129.241.187.255", PORT}
	broadcastUDP, _ := net.ResolveUDPAddr("udp", strings.Join(broadcastAddr, ""))
	broadcastConn, _ := net.DialUDP("udp", nil, broadcastUDP)
	defer broadcastConn.Close()
	for {
		buf, err := json.Marshal(<-channel)
		if err == nil {
			broadcastConn.Write(buf)
		}
	}
}

func UDPListener(channel chan def.Message) {
	UDPReceiveAddr, err := net.ResolveUDPAddr("udp", PORT)
	if err != nil {
		fmt.Println(err)
	}

	UDPConn, err := net.ListenUDP("udp", UDPReceiveAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer UDPConn.Close()

	buf := make([]byte, 2048)
	trimmed_buf := make([]byte, 1)
	var received_message def.Message

	for {
		n, _, _ := UDPConn.ReadFromUDP(buf)
		trimmed_buf = buf[:n]
		err := json.Unmarshal(trimmed_buf, &received_message)
		if err == nil {
			channel <- received_message
		}
	}
}
