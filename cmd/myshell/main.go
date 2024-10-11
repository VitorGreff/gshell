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
			fmt.Fprint(os.Stdout, "error while reading command")
			os.Exit(1)
		}

		shellString = shellString[:len(shellString)-1]
		args := strings.Split(shellString, " ")
		command := args[0]
		if !isWithinAvailableCommands(command) {
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
		executable := args[1]
		pathString := os.Getenv("PATH")
		if isWithinAvailableCommands(executable) {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", executable)
			return
		}
		isOnPath(pathString, executable)
	default:
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

func isWithinAvailableCommands(command string) bool {
	for _, c := range validCommands {
		if command == c {
			return true
		}
	}
	return false
}

func isOnPath(path string, exe string) {
	addresses := strings.Split(path, ":")
	for _, address := range addresses {
		files, err := os.ReadDir(address)
		if err != nil {
			continue
		}
		for _, file := range files {
			if exe == file.Name() {
				fmt.Fprintf(os.Stdout, "%s is %s/%s\n", exe, address, exe)
				return
			}
		}
	}
	fmt.Fprintf(os.Stdout, "%s: not found\n", exe)
}
