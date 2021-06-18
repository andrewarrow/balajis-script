package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("./balajis [directory]")
		return
	}
	PrintEveryClout(os.Args[1])
}
