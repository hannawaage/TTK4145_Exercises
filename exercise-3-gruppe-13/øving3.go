package main

import (
    "net"
	"fmt"
	"time"
)



func TCP(){
	fmt.Println("Hei")
	buf:= make([]byte, 1024)

	ServerAddr1, err := net.ResolveTCPAddr("tcp","10.100.23.147:33546")
	conn1, err := net.DialTCP("tcp", nil, ServerAddr1)

	if err != nil {
		fmt.Println("error 1")
		return 
	}


	fmt.Println("error 2")

	msg := "Connect to: 10.100.23.210:20013\x00"


	ServerAddr2, err := net.ResolveTCPAddr("tcp","10.100.23.210:20013")
	listener2, err := net.ListenTCP("tcp", ServerAddr2)

	fmt.Println("error 3")

	_, err = conn1.Write([]byte(msg))
	if err != nil {
		fmt.Println("Error sending")
		return 
	}

	fmt.Println("Hei")

	conn2, err := listener2.AcceptTCP()


	fmt.Println("Hei")

	if err != nil {
		fmt.Println("error 4")
		return 
	}


	
	fmt.Println("error 5")
	defer conn2.Close()
	i := 0
	for  {
		conn2.Write([]byte("Hello\x00"))
		len, err := conn2.Read(buf)
		//fmt.Println("Hei forløkke")
		if err != nil {
			return 
		}

		fmt.Println(string(buf[0:len]))
		i += 1
		time.Sleep(100 * time.Millisecond)
	}
}

func UDP(){

	buf := make([]byte, 1024)

	//ServerAddr1, err := net.ResolveUDPAddr("udp","10.100.23.147:20013")
	conn1, err := net.Dial("udp", "10.100.23.147:20013")

	if err != nil {
		return 
	}

	ServerAddr2, err := net.ResolveUDPAddr("udp","10.100.23.210:20013")
	
	conn2, err := net.ListenUDP("udp", ServerAddr2)

	if err != nil {
		return 
	}
	defer conn2.Close()
	i := 0
	for i < 20 {
		conn1.Write([]byte("Ping"))
		len, err := conn2.Read(buf)

		if err != nil {
			return 
		}

		fmt.Println(string(buf[0:len]))
		i += 1
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	UDP()
	TCP()

}
