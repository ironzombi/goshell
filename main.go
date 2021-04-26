package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fsec/goshell/cmd"
)

func listCwd() {
	var files []string
	root, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}

func currentDir() string {
	cdir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cdir)
	return cdir
}

func echoCmd(carg []string) {
	fmt.Println(carg[1:])
}

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
			listCwd()
		case command[0] == "pwd":
			currentDir()
		case command[0] == "echo":
			echoCmd(command)
		case command[0] == "hello":
			cmd.ListAll()
		case command[0] == "clear":
			cmd.ClearScreen()
		default:
			fmt.Println("command not found")
		}
	}

}
