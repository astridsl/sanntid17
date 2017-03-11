package src

import {
	"./Network"
	"./config"
}

func main() {
	go Network.main()//Setter i gang network og sjekker etter og bestemmer hele tiden  master

	ButtonEvents := make(chan config.buttonEvent) //Channel som sender external order til nettverk 
	OrderCompleted := make(chan bool)
	ArrivedAtFloor := make(chan bool)
	NewOrder := make(chan bool) 

	//Cost




	/*for {
		if Network.isMaster {
			//Heis er master --> deleger ordre

		} else {
			//Heis er slave --> do something

		}
	}*/

	//HUSK Å INITIALISERE driver MED elev.init --> LAG ERROR-FUNKSJON HVIS DET SKJER NOE FEIL

}