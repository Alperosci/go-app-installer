package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`###Should used with sudo###
Made by Alperosci - v1
Usage installexecutable [AppName]`)
	} else if os.Getuid() != 0 {
		fmt.Println("Run with sudo !!")
	} else {
		exec.Command("mv", os.Args[1], "/bin").Run()
	}
}
