package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var validCommands = []string{
	"exit", "echo", "type",
}

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
		if !isValid(command) {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		} else {
			evaluateCommand(command, args)
		}
	}
}

func evaluateCommand(command string, args []string) {
	switch command {
	case "exit":
		if args[1] != "0" {
			os.Exit(1)
		}
		os.Exit(0)
	case "echo":
		echoString := strings.Join(args[1:], " ")
		fmt.Fprintf(os.Stdout, "%s\n", echoString)
	case "type":
		typeShell := args[1]
		if isValid(typeShell) {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", typeShell)
		} else {
			fmt.Fprintf(os.Stdout, "%s: not found\n", typeShell)
		}
	default:
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

func isValid(command string) bool {
	for _, c := range validCommands {
		if command == c {
			return true
		}
	}
	return false
}
