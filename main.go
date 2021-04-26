package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"fsec/goshell/cmd"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[H\033[2J")
	fmt.Println("FSEC::SHELL")
	fmt.Println("___________")
	fmt.Print("#:/ ")

	for {
		fmt.Print(":> ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)
		command := strings.Fields(text)

		switch {
		case command[0] == "exit":
			return
		case command[0] == "uname":
			fmt.Println("goshell v0.1")
		case command[0] == "ls":
			cmd.ListCwd()
		case command[0] == "pwd":
			cmd.CurrentDir()
		case command[0] == "echo":
			cmd.EchoCmd(command)
		case command[0] == "hello":
			cmd.ListAll()
		case command[0] == "clear":
			cmd.ClearScreen()
		default:
			fmt.Println("command not found")
		}
	}

}
