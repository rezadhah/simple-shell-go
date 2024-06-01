package command

import (
	"github.com/rezadhah/shell/internal/errors"
)

func Factory(cmd string) (Command, error) {
	switch cmd {
	case COMMAND_CD:
		return NewCDCommand(), nil
	case COMMAND_EXIT:
		return NewExitCommand(), nil
	default:
		return nil, errors.ErrInvalidCmd
	}
}
