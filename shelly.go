package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		//Read the keyboard input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		//Handle the execution of the input
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	//Remove the newline character
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	//Check in for built-in commands
	switch args[0] {
	case "cd":
		//'cd' to home dir with empty path not yet supported
		if len(args) < 2 {
			return errors.New("path required")
		}
		//cange the directory and return the error
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}
	//prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...)

	//Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//exec the command and return the error
	return cmd.Run()

}