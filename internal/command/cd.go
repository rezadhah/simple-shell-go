package command

import (
	"os"

	"github.com/rezadhah/shell/internal/errors"
)

type cd struct{}

func NewCDCommand() Command {
	return &cd{}
}

func (c *cd) Execute(args ...string) error {
	if len(args) < 2 {
		return errors.ErrNoPath
	}
	return os.Chdir(args[1])
}
