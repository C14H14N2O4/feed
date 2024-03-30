package listen

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
)

func Listen() (bool, error) {
	port := os.Args[2]
	// check if given port is a number
	if _, err := strconv.Atoi(port); err != nil {
		return false, errors.New("port provided is not a valid integer")
	} else {
		fmt.Println("Attempting to listen on port: ", port)
	}

	// verify if port is open
	// port 135 is open
	if portStatus(port) {
		fmt.Println(port + " is open")
		return true, nil
	} else {
		return true, errors.New("Port " + port + " is not open")
	}

	// if port is open,
}

func portStatus(port string) bool {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		return false
	}
	// defer executes when surrounding function returns
	defer conn.Close()
	return true
}
