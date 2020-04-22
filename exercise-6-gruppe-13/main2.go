package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
	t "time"

	//"encoding/json"
	"log"
	"os/exec"
)

var counter uint64
var buf = make([]byte, 16)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		log.Fatal(err)
	}
}

func main() {

	//Become backup, and listen
	addr1, err := net.ResolveUDPAddr("udp", "localhost:20013")
	master := false

	conn1, err := net.ListenUDP("udp", addr1)
	checkError(err)

	fmt.Println("For now, I am backup")

	for !(master) {
		conn1.SetReadDeadline(t.Now().Add(3 * t.Second))
		len, _, err := conn1.ReadFromUDP(buf)
		if err != nil {
			master = true
		} else {
			counter = binary.BigEndian.Uint64(buf[:len])
		}
	}
	//No new message, I must become master

	conn1.Close()

	//make new backup
	backup := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run main2.go")
	backup.Run()
	checkError(err)

	fmt.Println("Now, I am master")

	conn2, err := net.Dial("udp", "localhost:20013")
	checkError(err)

	for {
		fmt.Println(counter)
		counter++
		time.Sleep(200 * time.Millisecond)
		binary.BigEndian.PutUint64(buf, counter)
		_, _ = conn2.Write(buf)
	}
}
