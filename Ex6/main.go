package main

import (
	"./network/bcast"
  "./network/localip"
  "./network/peers"
  "fmt"
  "runtime"
	"io/ioutil"
	"os"
)
const filename := "backupfile.txt"

func main() {
	var id string
	if id == "" {
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		id = fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())
	}

	primaryPing := make(chan)
	primaryDown := make(chan)
	go peers.Receiver(15647, primaryPing)
	go primaryCkecker(primaryPing, primaryDown)

	<-primaryDown //wait for primary to go down

	file, err := os.Open(filename, os.O_CREATE, 0755) // For read access.
	info, err := os.Stat(filename)
	if info.Size() == 0 {
		file.write(1)
	}
	file.close();
	go peers.Transmitter(15647, id, primaryPing)
	pinger(primaryPing)

	for now := range time.Tick(0.5*time.Second) {
		file, err := os.Open(filename) // For read access.
		fmt.Println(file.read())
		file.close()
	}
}

func primaryCkecker(primaryPing chan, primaryDown chan) {
	timeout := make(chan timer)
	for {
		timeout = time.NewTimer(200 * time.MilliSecond)

		select {
		case: primaryPing
		// do nothing
		case: timeout
	  primaryDown <-
		break; //Break loop and end
		}
	}
}

func pinger(primaryPing chan) {
	for now := range time.Tick(20 * time.MilliSecond) {
		primaryPing <-
	}
}
