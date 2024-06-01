package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	errors "github.com/rezadhah/shell/internal/errors"
)

const (
	CD_COMMAND   = "cd"
	EXIT_COMMAND = "exit"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case CD_COMMAND:
		if len(args) < 2 {
			return errors.ErrNoPath
		}

		return os.Chdir(args[1])
	case EXIT_COMMAND:
		os.Exit(0)
	default:
		return runCommandUsingBinary(args...)
	}

	return nil
}

func runCommandUsingBinary(args ...string) error {
	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}
