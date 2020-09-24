package main

import (
	"fmt"
	"os"
)

func getWorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return wd
}

func main() {
	fmt.Println("Cruel Test for GoLang")
	wd := getWorkingDirectory()
	fmt.Println("Working Directory is " + wd)
}
