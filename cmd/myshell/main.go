package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		// Wait for user input
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprint(os.Stdout, "Error while reading command")
			os.Exit(1)
		}
		command = command[:len(command)-1]
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}
