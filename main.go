package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
)

type Config struct {

	// Thanks. I hate it.
	Hosts struct {
		Host []struct {
			Address        string `yaml:"address"`
			Port           string `yaml:"port"`
			File_Directory string `yaml:"known_hosts_file"`
		} `yaml:"server"`
	} `yaml:"servers"`
}

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

	// Fuck spaces.
	known_hosts_str := strings.Split(dialAddr, ":")[0] + " " + key.Type() + " " + base64.StdEncoding.EncodeToString(key.Marshal()) + "\n"

	fmt.Printf("\nFingerprint of host: %s\n\n%s %s\n\n", dialAddr, key.Type(), base64.StdEncoding.EncodeToString(key.Marshal()))
	fmt.Println("Checking if already exists in file before writing")

	check := ReadFile(known_hosts_str, file_directory)

	if !check { // Only write the key to the file if false

		log.Println("Writing to file")
		FileOp(known_hosts_str, file_directory)
	} else {

		log.Printf("Key already exist in known_hosts file")
	}

	return nil
}

func ReadFile(data string, data_file string) bool {
	contents, _ := os.ReadFile(data_file)

	if strings.Contains(string(contents), data) {
		return true
	} else {
		return false
	}
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

	file.Close()

}

func ChkYaml(file *string) Config {

	var config Config

	_, err := os.Stat(*file) // check if config exists

	if err == nil { // If exists, read file
		data, err := os.ReadFile(*file)
		if err != nil { // Check for any errors reading the file
			log.Fatal(err)
		}

		if err := yaml.Unmarshal(data, &config); err != nil { // Try to unmarshal. Check for any errors.
			log.Fatal(err)
		}
	}

	return config

}

func DialSSH(host_str string, port_str int) {
	sshConfig := &ssh.ClientConfig{
		HostKeyCallback: KeyPrint,
	}

	ssh.Dial("tcp", fmt.Sprintf("%s:%d", host_str, port_str), sshConfig)

}

func main() {

	var config Config
	var config_set bool

	host_file := flag.String("c", "", "The yaml config file with addresses and ports of hosts.")
	host := flag.String("h", "", "The SSH host to fingerprint.")
	port := flag.Int("p", 22, "The SSH port of host.")
	file := flag.String("f", ".ssh/known_hosts", "Known hosts file location.")
	flag.Parse()

	// Check if either config file was passed or host flag was passed

	if !(isFlagPassed("h") || isFlagPassed("c")) {
		usage()
		log.Fatal("\nNot enough arguments passed.")
	}

	if isFlagPassed("c") {

		config = ChkYaml(host_file)
		config_set = true
	}

	var host_str string
	var port_str int

	if config_set {
		for _, host := range config.Hosts.Host {

			if host.Address == "" {
				log.Fatal("Please specify host of the system you want to fingerprint with ssh")
			} else {
				host_str = host.Address
			}

			// Check if the port of the host was set.
			if host.Port == "" { // If empty or not set, set to the default 22 port for ssh
				port_str = 22
			} else {
				// Convert the string to int and assign to the variable
				port, _ := strconv.Atoi(host.Port)
				port_str = port
			}

			if host.File_Directory == "" {
				home, _ := os.UserHomeDir()
				file_directory = home + `/.ssh/known_hosts`
			} else {
				file_directory = host.File_Directory
			}

			DialSSH(host_str, port_str)
		}

	} else {

		/*
			Assign the flag pointer variables to string variables.
			Probably a better way, but I really don't care and
			don't want to put in more effort to possibly overcomplicate
			something when I could do it this way.
		*/

		file_directory = *file

		host_str = *host
		port_str = *port

		DialSSH(host_str, port_str)
	}

}
