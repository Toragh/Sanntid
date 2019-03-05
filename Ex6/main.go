package main

import (
	"./network/bcast"
	"./network/localip"
	"fmt"
	"os"
	"time"
"encoding/binary"
"log"
"bytes"
"os/exec"
)

type database struct {
	Data uint32
}

func main() {
	var id string
	filename := "backup.bin"
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
//Backup ends
	<-primaryDown //wait for primary to go down
//Primary starts
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}

		s := &database{
					0,
				}
		var bin_buf bytes.Buffer
		binary.Write(&bin_buf, binary.BigEndian, s)
		writeNextBytes(file, bin_buf.Bytes())
	  file.Close()

	}


	go bcast.Transmitter(15647, primaryPing)
	go pinger(primaryPing)

	cmd := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run main.go")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	for range time.Tick(500*time.Millisecond) {
		file, err := os.OpenFile(filename, os.O_RDWR, 0755)
		m := database{}
		data := readNextBytes(file, 16)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &m)
		if err != nil {
			log.Fatal("binary.Read failed", err)
		}
		file.Close()
		m.Data +=1;
		fmt.Println(m.Data)

		file, err = os.OpenFile(filename, os.O_WRONLY, 0755)
		var bin_buf1 bytes.Buffer
		binary.Write(&bin_buf1, binary.BigEndian, m)
		writeNextBytes(file, bin_buf1.Bytes())
	  file.Close()
	}
}


func writeNextBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
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
