package listen

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Listen() (bool, error) {
	// verify port is number; check to see if it is open
	port := os.Args[2]
	if _, err := strconv.Atoi(port); err != nil {
		return false, errors.New("port provided is not a valid integer")
	} else {
		fmt.Println("Listening on port: ", port)
		return true, nil
	}
}
