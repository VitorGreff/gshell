package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	command = command[:len(command)-1]
	fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
}
