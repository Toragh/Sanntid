package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func receive() {
	/* Lets prepare a address at any address at port 30000*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":30000")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received @@", string(buf[0:n]), "@@ from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func sendAndReceive() {
	buf := make([]byte, 1024)

	ServerAddr, err := net.ResolveUDPAddr("udp", "10.100.23.242:20016")
	CheckError(err)

	Conn, err := net.DialUDP("udp", nil, ServerAddr)
	CheckError(err)
	defer Conn.Close()

	LocalAddr, err := net.ResolveUDPAddr("udp", ":20016")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", LocalAddr)
	CheckError(err)
	defer ServerConn.Close()

	fmt.Println("End of init")
	defer Conn.Close()
	for {
		fmt.Println("Foor loop")
		msg := "Hei paa deg"
		_, err := Conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(msg, err)
		}

		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received @@", string(buf[0:n]), "@@ from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {
	sendAndReceive()
}
