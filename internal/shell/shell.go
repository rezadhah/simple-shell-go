package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rezadhah/shell/internal/command"
	"github.com/rezadhah/shell/internal/errors"
)

type Store struct {
	allowedCommands map[string]struct{}
}

func NewShell() *Store {
	store := &Store{
		allowedCommands: make(map[string]struct{}),
	}

	store.init()
	return store
}

func (s *Store) init() {
	for _, cmd := range command.AllowedCommands() {
		s.allowedCommands[cmd] = struct{}{}
	}
}

func (s *Store) InputMode() error {
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

func (s *Store) GetAllowedCommands() []string {
	cmds := make([]string, 0)
	for k, _ := range s.allowedCommands {
		cmds = append(cmds, k)
	}
	return cmds
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case command.COMMAND_CD:
		if len(args) < 2 {
			return errors.ErrNoPath
		}

		return os.Chdir(args[1])
	case command.COMMAND_EXIT:
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
