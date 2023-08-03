package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s: \n", os.Args[0])
	fmt.Println()
	flag.PrintDefaults()
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

var file_directory string // Don't ask.

func KeyPrint(dialAddr string, addr net.Addr, key ssh.PublicKey) error {
	fmt.Printf("Fingerprint of host: %s\n\n%s %s\n\n", dialAddr, key.Type(), base64.StdEncoding.EncodeToString(key.Marshal()))
	log.Println("Writing to file")

	// Fuck spaces.
	FileOp(strings.Split(dialAddr, ":")[0]+" "+key.Type()+" "+base64.StdEncoding.EncodeToString(key.Marshal())+"\n", file_directory)
	return nil
}

func FileOp(data string, data_file string) { // I haven't slept.

	file, err := os.OpenFile(data_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Create and open file to append key to

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil { // Write data to file
		log.Fatal(err)
	}

}

func main() {
	host := flag.String("h", "", "The SSH host to fingerprint.")
	port := flag.Int("p", 22, "The SSH port of host.")
	file := flag.String("f", ".ssh/known_hosts", "Known hosts file location.")
	flag.Parse()

	if !isFlagPassed("h") {
		usage()
		log.Fatal("\nNot enough arguments passed.")
	}

	file_directory = *file

	var host_str string = *host
	var port_str int = *port

	sshConfig := &ssh.ClientConfig{
		HostKeyCallback: KeyPrint,
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host_str, port_str), sshConfig)
	if err != nil {
		os.Exit(0)
	}
	defer client.Close()
}
