package listen

import (
	"fmt"
	"os"
)

func Listen() {
	fmt.Println("Args:")
	for i, arg := range os.Args {
		fmt.Printf("Argument %d: %s\n", i, arg)
	}
}
