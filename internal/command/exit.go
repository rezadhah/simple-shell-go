package command

import (
	"os"

	"github.com/rezadhah/shell/internal/errors"
)

type exit struct{}

func NewExitCommand() Command {
	return &exit{}
}

func (e *exit) Execute(args ...string) error {
	if len(args) != 1 {
		return errors.ErrInvalidArgs
	}
	return os.Chdir(args[0])
}
