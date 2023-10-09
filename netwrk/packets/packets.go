package main

import (
	"fmt"
	"log"
	"net"
	"os"

	// sudo apt-get install libpcap-dev before
	"github.com/google/gopacket/pcap"
)

func checkDeviceExist(dev string) bool {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	for _, dv := range devices {
		if dv.Name == dev {
			return true
		}
	}

	return false
}

func enumerateDevices() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range devices {
		fmt.Printf("name - %v\n", device.Name)
		fmt.Printf("description - %v\n", device.Description)
		fmt.Printf("adress - %v\n", device.Addresses)
		fmt.Println()
	}
}

func main() {
	fmt.Printf("ens33 exist - %v\n", checkDeviceExist("ens33"))

	strEcho := "hallo\n"
	servAddr := "localhost:8081"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}
}
