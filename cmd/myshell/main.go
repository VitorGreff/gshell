package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var validCommands = []string{
	"exit", "echo", "type", "pwd", "cd",
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		shellString, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprint(os.Stdout, "error while reading command\n")
			os.Exit(1)
		}
		// remove \n
		shellString = shellString[:len(shellString)-1]
		args := strings.Split(shellString, " ")
		command := args[0]
		evaluateCommand(command, args)
	}
}

func evaluateCommand(command string, args []string) {
	pathString := os.Getenv("PATH")
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
		fileName := args[1]
		if isWithinAvailableCommands(fileName) {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", fileName)
			return
		}
		flag, fileLocation := isWithinPath(pathString, fileName)
		if flag {
			fmt.Fprintf(os.Stdout, "%s is %s\n", fileName, fileLocation)
			return
		}
		fmt.Fprintf(os.Stdout, "%s: not found\n", fileName)
	case "pwd":
		absolutePath, _ := os.Getwd()
		fmt.Fprintln(os.Stdout, absolutePath)
	case "cd":
		path := args[1]
		if path == "~" {
			path, _ = os.UserHomeDir()
		}
		err := os.Chdir(path)
		if err != nil {
			fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", path)
		}
	default:
		flag, _ := isWithinPath(pathString, command)
		if flag {
			cmd := exec.Command(command, args[1:]...)
			log, _ := cmd.Output()
			// remove \r\n
			fmt.Fprintln(os.Stdout, string(log[:len(log)-1]))
			return
		}
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

func isWithinPath(path string, givenFile string) (bool, string) {
	addresses := strings.Split(path, ":")
	for _, address := range addresses {
		files, err := os.ReadDir(address)
		if err != nil {
			continue
		}
		for _, file := range files {
			if givenFile == file.Name() {
				return true, fmt.Sprintf("%s/%s", address, givenFile)
			}
		}
	}
	return false, ""
}
