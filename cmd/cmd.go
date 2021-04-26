package cmd

import (
	"fmt"
)

func ListAll() {
	fmt.Println("hey")
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
