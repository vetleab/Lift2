package main

import (
	"driver"
	"network"
	"operations"
	"fmt"
	//"os/exec"
	"time"
)

func Initialize() bool {
	driver.Elev_init()
	driver.Elev_set_motor_direction(-1)
	for {
		if driver.Elev_get_floor_sensor_signal() != -1 {
			driver.Elev_set_motor_direction(0)
			driver.Elev_set_floor_indicator(driver.Elev_get_floor_sensor_signal())
			return true
		}
	}
}

func main() {
	if Initialize() {
		fmt.Printf("Started!\n")
	} else {
		fmt.Printf("error!\n")
	}

	operations.Laddr = "129.241.187.158"

	var outgoingMsg = make(chan operations.Udp_message, 10)
	var incomingMsg = make(chan operations.Udp_message, 10)

	go operations.Request_buttons()
	go operations.Request_floorSensor()
	//go operations.Fsm_printstatus()
	go operations.Request_timecheck()

	object := operations.Udp_message{Category: 1, Floor: 1, Button: 1, Cost: 1}
	network.Init(outgoingMsg, incomingMsg)	
	
	fmt.Println("PIKK")
	for{
		outgoingMsg <- object
		object.Floor = object.Floor + 1
		//fmt.Println("Sending object")
		time.Sleep(time.Second * 2)
	}

	//operations.CloseConnectionChan <- true


	//cmd := exec.Command("gnome-terminal", "-x", "go", "run", "main.go")
	//cmd.Run()
	
	//msg := network.Udp_message{Raddr: "broadcast", Data: object.Data, Length: 0}

}
