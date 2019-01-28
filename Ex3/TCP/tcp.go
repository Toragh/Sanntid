package main

import (
	"fmt"
	"log"
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

//forferdelig funksjonsnavn
func ReadFromWriteTo(socket *net.TCPConn, doneChannel chan bool) {
	var buffer [1024]byte
	var msg = "I am the tuna\x00"
	for {
		//read from server
		_, err := socket.Read(buffer[:])
		CheckError(err)
		log.Println(string(buffer[:]))

		//write to server
		_, err = socket.Write([]byte(msg))
		CheckError(err)
		//delay
		time.Sleep(2000 * time.Millisecond)
	}
	doneChannel <- true
}

//forferdelig funksjonsnavn
func accept_connection(socketConnect *net.TCPConn, doneChannel chan bool) {
	//new message
	var buffer [1024]byte
	var msg = "I am the crab\x00"
	for {
		//read from server
		_, err := socketConnect.Read(buffer[:])
		CheckError(err)
		log.Println(string(buffer[:]))

		//write to server
		_, err = socketConnect.Write([]byte(msg))
		CheckError(err)

		//delay
		time.Sleep(2000 * time.Millisecond)
	}
	doneChannel <- true
}

func main() {
	//TCP-setup
	raddr, err := net.ResolveTCPAddr("tcp", "10.100.23.242:34933")
	CheckError(err)

	socket, err := net.DialTCP("tcp", nil, raddr)
	CheckError(err)

	//declare variables
	var msg = "Connect to: 10.100.23.138:20016\x00"

	//make listener
	laddr, err := net.ResolveTCPAddr("tcp", ":20016")
	CheckError(err)

	listener, err := net.ListenTCP("tcp", laddr)
	CheckError(err)

	//server connect back
	_, err = socket.Write([]byte(msg))
	CheckError(err)

	socketConnect, err := listener.AcceptTCP()
	CheckError(err)

	doneChannel := make(chan bool, 1)

	go ReadFromWriteTo(socket, doneChannel)
	go accept_connection(socketConnect, doneChannel)

	<-doneChannel
}
