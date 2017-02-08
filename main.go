package main

import (
	"./Interface/driver"
	"./Network"
	"C"
	"fmt"
)

func main() {
	fmt.Println("Hello Github!")

	C.elev.elev_init()

	Netmain.Netmain()
}
