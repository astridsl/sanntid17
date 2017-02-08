package main

import (
	"./Interface/driver"
	//"./Network"
	"fmt"
)

func main() {
	fmt.Println("Hello Github!")

	driver.Elev_init()

	driver.Elev_set_motor_direction(-1)
	//Netmain.Netmain()
}
