package main

import (
	"./network/bcast"
	"./network/localip"
	"fmt"
	"os"
	"time"
"io/ioutil"
"encoding/binary"
)



func main() {
	var id string
	filename := "backupfile.txt"
	if id == "" {
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		id = fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())
	}
	primaryPing := make(chan int)
	primaryDown := make(chan int)
	go bcast.Receiver(15647, primaryPing)
	go primaryCkecker(primaryPing, primaryDown)

	<-primaryDown //wait for primary to go down
	
	ioutil.WriteFile("/tmp/dat1", d1, 0644)
	info, err := os.Stat(filename)
	if info.Size() == 0 {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
		}
		file.WriteString(fmt.Sprintf("%d\n", 0))
		file.Close();
	}

	go bcast.Transmitter(15647, id, primaryPing)
	pinger(primaryPing)
	
	for range time.Tick(500*time.Millisecond) {
		file, err := os.Open(filename) // For read access.
		if err != nil {
			fmt.Println(err)
		}
		b := make([]byte, 4)
    		file.Read(b)
		i := binary.BigEndian.Uint32(b)
		fmt.Println(i)
		i += 1
		file.WriteString(fmt.Sprintf("%d\n", i))
		file.Close()
	}
}

func primaryCkecker(primaryPing chan int, primaryDown chan int) {
	for {
		select {
		case <-primaryPing:
			// do nothing
		case <-time.After(300 * time.Millisecond):
	  		primaryDown <- 1
			return //Break loop and end
		}
	}
}

func pinger(primaryPing chan int) {
	for range time.Tick(100 * time.Millisecond) {
		primaryPing <- 1
	}
}

