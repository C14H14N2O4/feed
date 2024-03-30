package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument")
	} else {
		switch os.Args[1] {
		case "sitrep": // shows current chats
			fmt.Println("sitrep")
		case "burn": // clears all connections & all logs if implemented // -perma flag to add ips to ban list
			fmt.Println("burn")
		case "revoke": // revoke <<uid>> <<opt:PERMA>>: terminate connection, if perma flag given, add ip to ban list
			fmt.Println("revoke")
		case "flush": // flush <<uid>>: clears logs relating to specific chat (if implemented)
			fmt.Println("flush")
		case "ping": // ping <<uid OR addr, port>>: checks if uid/ addr and port listening
			fmt.Println("ping")
		case "connect":
			fmt.Println("connect")
		case "status": // status <<uid>>: checks if chat is connected/pending/rejected
			fmt.Println("status")
		case "listen": // listen <<port>>: begins to listen at port <<port>> for incoming connections
			fmt.Println("listen")
		case "quarantine": // quarantine <<list>>: shows all connection requests // // quarantine <<purge, opt: PERMA>>: deletes all connection requests, optionally adds all ips to block list
			fmt.Println("quarantine")
		case "help": // prints help info
			fmt.Println("help")
		default:
			fmt.Printf("Command not recognized, please run feed help to view full list of commands")
		}
	}

}
