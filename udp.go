package main

import (
	"fmt"
	"log"
	"os"
	"net"
	"time"
)

func main() {
	// listen to incoming udp packets
	udpServer, err := net.ListenPacket("udp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()

	for {
		buf := make([]byte, 10*1024)
		_, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			continue
		}
		go response(udpServer, addr, buf)
	}

}

func response(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("time received: %v. Your message: %v!", time, getFiles(string(buf)))
	fmt.Println(responseStr)
	udpServer.WriteTo([]byte(responseStr), addr)
}


func getFiles(path string) {
	fileStat, err := os.Stat(path)
 
	if err != nil {
		log.Fatal(err)
	}
 
	fmt.Println("File Name:", fileStat.Name())        // Base name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", fileStat.IsDir())   // Abbreviation for Mode().IsDir()
}