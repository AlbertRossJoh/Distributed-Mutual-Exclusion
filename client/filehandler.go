package main

import (
	"log"
	"os"
	"strings"
)

func WriteToSharedFile() {
	log.Printf("The client's ip address is: ", clientIpAddr)

	f, err := os.OpenFile("/var/nas/shared", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Couldn't open the shared file! On client ", clientIpAddr)
	}
	defer f.Close()
	if _, err := f.WriteString(clientIpAddr + "\n"); err != nil {
		log.Println("Couldn't write to shared file! On client ", clientIpAddr)
	}
}

func GetFileContents() []string {
	f, err := os.OpenFile("/var/nas/shared", os.O_RDONLY, 0644)
	if err != nil {
		log.Println("Couldn't open the shared file! On client ", clientIpAddr)
	}
	defer f.Close()
	bytes := make([]byte, 1_000_000)
	f.Read(bytes)
	contents := string(bytes)
	strs := strings.Split(contents, "\n")
	if len(strs) > 0 {
		// Removes whitespace from end of file
		strs = strs[:len(strs)-1]
	}
	return strs
}
