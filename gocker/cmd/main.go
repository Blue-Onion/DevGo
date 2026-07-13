package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		fmt.Println("What the fok")
	}

}
func parent() {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(exe, append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}
}
func child() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}
}
func must(err error) {
	if err != nil {
		panic(err)
	}
}
