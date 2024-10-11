package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdExit string = "exit"
)

func main() {
	for {
		// Wait for user input
		fmt.Fprint(os.Stdout, "$ ")
		shellString, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprint(os.Stdout, "Error while reading command")
			os.Exit(1)
		}

		shellString = shellString[:len(shellString)-1]
		args := strings.Split(shellString, " ")
		command := args[0]

		switch command {
		case CmdExit:
			if args[1] != "0" {
				os.Exit(1)
			}
			os.Exit(0)
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		}
	}
}
