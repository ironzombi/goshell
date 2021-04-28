package cmd

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"
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

// Ping {host} well...connect on tcp22
func Ping(host string) {
	port := "22"
	timeout := time.Duration(1 * time.Second)
	_, err := net.DialTimeout("tcp", host+":"+port, timeout)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s %s %s\n", host, "responding on port:", port)
	}
}

// write * to * echo to a file
func WriteFile(file string, data string) {
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(fmt.Sprintf("%s", data[0:]))
	if err2 != nil {
		fmt.Println(err)
	}
}
