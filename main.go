package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
		}
		input = strings.TrimSuffix(input, "\n")
		if input == "" {
		continue
		}

		if err := executeInput(input); err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}

var ErrNoPath = errors.New("no path provided")

func executeInput(input string) error {
	args := strings.Split(input," ")
	switch args[0]{
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}
		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}


