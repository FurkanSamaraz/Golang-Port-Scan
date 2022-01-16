package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {

	gorsel := `
                      __                      	
    ____  ____  _____/ /_______________ _____ 
   / __ \/ __ \/ ___/ __/ ___/ ___/ __ \/ __ \
  / /_/ / /_/ / /  / /_(__  ) /__/ /_/ / / / /
 / .___/\____/_/   \__/____/\___/\__,_/_/ /_/ 
/_/                                           
                                      ©Furkan Samaraz`

	fmt.Println(gorsel)
	fmt.Println("Port Tarama")
	for i := 130; i <= 140; i++ {
		open := scanPort("tcp", "google.com", i)
		println(i, "Portlar :", open)
	}

}
func scanPort(protokol, hostname string, port int) bool {

	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protokol, address, 1*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
