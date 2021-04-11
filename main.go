package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

		switch {
		case text == "exit":
			return
		case text == "wow":
			fmt.Println("switch works!")
		case text == "ls":
			listCwd()
		case text == "pwd":
			currentDir()
		default:
			fmt.Println("command not found")
		}
	}

}
