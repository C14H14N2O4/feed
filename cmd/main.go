package main

import (
	"feed/internal/pkg/listen"
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
		case "burn": // clears all connections & all logs if implemented
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
			// check to see if port is provided -> if yes then if port is open -> then start subroutine at port
			if len(os.Args) < 3 {
				fmt.Println("Please provide a port")
			} else {
				listen.Listen()
				// if reflect.TypeOf(os.Args[2]).String() != "string" {
				// 	fmt.Println("Invalid port selection")
				// } else {
				// 	fmt.Println("Listening on port " + os.Args[2])
				// }
			}
		case "mute": // mute <<port>>: kills the listen process at port <<port>>
			fmt.Println("mute")
		case "quarentine": // quarentine <<list>>: shows all connection requests // // quarentine <<purge, opt: PERMA>>: deletes all connection requests, optionally adds all ips to block list
		case "help": // prints help info
			fmt.Println("help")
		case "message": // message <<uid>> attempts to send a message to addr, port referenced by <<uid>>
			fmt.Println("message")
		default:
			fmt.Printf("Command not recognized, please run feed help to view full list of commands")
		}
	}

}
