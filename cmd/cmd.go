package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

//  ListAll is a sleeper function
func ListAll() {
	fmt.Println("hey")
}

// ClearScreen clears the current the terminal
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// list the current dir - go pid
func ListCwd() {
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
	fmt.Println(root)
}

// list the current dir path
func CurrentDir() string {
	cdir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cdir)
	return cdir
}

// echo echo
func EchoCmd(carg []string) {
	fmt.Println(carg[1:])
}
